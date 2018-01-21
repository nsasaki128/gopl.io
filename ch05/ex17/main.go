package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main url tags...")
	}

	url := os.Args[1]
	tags := os.Args[2:]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := ElementsByTagName(doc, tags...)

	for _, r := range result {
		fmt.Printf("Data: %v, Attr: %v\n", r.Data, r.Attr)
	}

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	set := map[string]bool{}

	for _, n := range name {
		set[n] = true
	}

	var find func(*html.Node, []*html.Node) []*html.Node

	find = func(n *html.Node, result []*html.Node) []*html.Node {
		if n.Type == html.ElementNode && set[n.Data] {
			result = append(result, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result = find(c, result)
		}
		return result
	}

	return find(doc, nil)
}
