package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	elements := make(map[string]int)
	outline(elements, doc)
	fmt.Println("element\tnum")
	for k, v := range elements {
		fmt.Fprintf(os.Stdout, "%s\t%d\n", k, v)
	}
}

func outline(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(elements, c)
	}
}
