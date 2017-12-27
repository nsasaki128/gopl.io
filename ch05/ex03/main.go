package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtexts: %v\n", err)
	}
	texts := findtexts(nil, doc)
	for _, v := range texts {
		fmt.Printf("%s\n", v)
	}
}

func findtexts(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = findtexts(texts, c)
	}
	return texts
}
