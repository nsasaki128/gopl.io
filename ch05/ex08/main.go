package main

import (
	"os"

	"net/http"

	"log"

	"fmt"

	"golang.org/x/net/html"
)

//forEachNOdeはnから始まるツリー内部の個々のノードxに対して
//関数pre(x)とpost(x)を呼び出します。その二つの関数はオプションです。
//preは子ノードを訪れた前に呼び出され
//postは子ノードを訪れた後に呼び出されます

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s url id\n", os.Args[0])
	}
	url, id := os.Args[1], os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	elem := ElementById(doc, id)
	if elem != nil {
		fmt.Println(elem)
	} else {
		fmt.Printf("element id %s is not found in %s\n", id, url)
	}

}

func ElementById(doc *html.Node, id string) *html.Node {

	var result *html.Node

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				result = n
				return false
			}
		}
		return true
	}
	post := func(n *html.Node) bool { return true }

	forEachNode(doc, pre, post)

	return result
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) (result bool)) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}
