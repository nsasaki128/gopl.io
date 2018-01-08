package main

import (
	"fmt"

	"os"

	"net/http"

	"golang.org/x/net/html"
)

//forEachNOdeはnから始まるツリー内部の個々のノードxに対して
//関数pre(x)とpost(x)を呼び出します。その二つの関数はオプションです。
//preは子ノードを訪れた前に呼び出され
//postは子ノードを訪れた後に呼び出されます

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
