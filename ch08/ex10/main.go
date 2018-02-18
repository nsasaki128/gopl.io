package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func crawl(url string) []string {
	if cancelled() {
		return nil
	}
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	worklist := make(chan []string)
	unseenLinks := make(chan string) //重複していないURL

	//コマンドラインの引数で開始する
	go func() { worklist <- os.Args[1:] }()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		go func() {
		loop:
			select {
			case link := <-unseenLinks:
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			case <-done:
				break loop
			}
		}()
	}

	//メインゴルーチンはworklistの項目の重複をなくし、
	//未探索の項目をクローラへ送る。
	seen := make(map[string]bool)
loop:
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					select {
					case unseenLinks <- link:
					case <-done:
						break loop
					}
				}
			}
		case <-done:
			break loop
		}
	}
}

//copy from ch05/link
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue //不正なURLを無視
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
