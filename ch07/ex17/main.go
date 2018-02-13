package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var id = flag.String("i", "", "search id")
var class = flag.String("c", "", "search class")

func main() {
	flag.Parse()
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement //要素のスタック
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) //プッシュ
		case xml.EndElement:
			stack = stack[:len(stack)-1] //ポップ
		case xml.CharData:
			if !containsAll(stackNames(stack), flag.Args()) {
				continue
			}
			if *id != "" || *class != "" {
				if len(stack) == 0 {
					continue
				}
				elem := stack[len(stack)-1]
				if *id != "" && !hasAttr(elem.Attr, "id", *id) {
					continue
				}
				if *class != "" && !hasAttr(elem.Attr, "class", *class) {
					continue
				}
			}
			fmt.Printf("%s: %s\n", strings.Join(stackNames(stack), " "), tok)

		}
	}
}

func stackNames(stacks []xml.StartElement) []string {
	var output []string
	for _, stack := range stacks {
		output = append(output, stack.Name.Local)
	}
	return output
}

func hasAttr(attr []xml.Attr, name, value string) bool {
	for _, a := range attr {
		if a.Name.Local == name && a.Value == value {
			return true
		}
	}
	return false
}

// containsAll はxがyの要素を順番に含んでいるかどうかを報告します。
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
