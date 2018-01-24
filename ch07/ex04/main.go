package main

import (
	"fmt"
	"io"
	"os"

	"log"

	"golang.org/x/net/html"
)

type myReader struct {
	s string
	i int64
}

func (r *myReader) Read(p []byte) (n int, err error) {
	//parse.goの1994行目を満たすための仕様
	// Iterate until EOF. Any other error will cause an early return.
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	//ここは単なるコピーを進めていくところ
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return n, nil
}

func newReader(s string) io.Reader {
	return &myReader{s, 0}
}
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: main url tags...")
	}
	r := newReader(os.Args[1])
	//parserの中でreaderを利用している
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
