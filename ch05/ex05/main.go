package main

import (
	"fmt"
	"net/http"

	"bufio"

	"strings"

	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s, s1", os.Args[0])
		os.Exit(1)
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("words:\t%d\nimages:\t%d\n", words, images)
}

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GET リクエストを urlへ
// 行い、そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cw, ci := countWordsAndImages(c)
		words += cw
		images += ci
	}

	return
}
