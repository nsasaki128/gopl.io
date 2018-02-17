package main

import (
	"fmt"
	"log"
	"net/http"

	"flag"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)

type work struct {
	links []string
	depth int
}

var depth = flag.Int("depth", 2, "depth for search")

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //トークンを獲得
	list, err := Extract(url)
	<-tokens //トークンを解放
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()
	worklist := make(chan *work)
	var n int //worklist への送信待ちの数
	//コマンドラインの引数で開始する
	n++
	go func() { worklist <- &work{links: flag.Args(), depth: 1} }()

	//ウェブを並行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		w := <-worklist
		if w.depth > *depth {
			continue
		}
		for _, link := range w.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- &work{links: crawl(link), depth: w.depth + 1}
				}(link)
			}
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
