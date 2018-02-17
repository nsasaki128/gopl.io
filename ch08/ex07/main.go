package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"io/ioutil"
	"net/url"
	"path/filepath"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //トークンを獲得
	list, err := Extract(url)
	<-tokens //トークンを解放
	if err != nil {
		log.Print(err)
	}
	filtered := selectSameHost(url, list)
	for _, link := range filtered {
		save(link)
	}
	return filtered
}

func main() {
	worklist := make(chan []string)
	var n int //worklist への送信待ちの数
	//コマンドラインの引数で開始する
	n++
	go func() { worklist <- os.Args[1:] }()

	//ウェブを並行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
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

//Copy from ch05/ex13
func selectSameHost(original string, list []string) []string {
	originalUrl, err := url.Parse(original)
	if err != nil {
		log.Println(err)
		return nil
	}

	var filtered []string

	for _, target := range list {
		compareUrl, err := url.Parse(target)
		if err != nil {
			log.Println(err)
			continue
		}
		if originalUrl.Host == compareUrl.Host {
			filtered = append(filtered, target)
		}
	}
	return filtered
}

func save(file string) {
	resp, err := http.Get(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	doc, err := url.Parse(file)
	if err != nil {
		log.Println(err)
		return
	}
	dir := filepath.Join("./", doc.Host, filepath.Clean(doc.Path))
	//for dividing dir path and file name
	base := "content"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(dir, base), bs, os.ModePerm)

	if err != nil {
		log.Println(err)
		return
	}
}
