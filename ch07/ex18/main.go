package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{} // CharDataあるいは*Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func NewTree(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	var result *Element
	result = nil
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			newElem := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				elem := stack[len(stack)-1]
				elem.Children = append(elem.Children, newElem)
			}
			stack = append(stack, newElem) //プッシュ
			if result == nil {
				result = newElem
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] //ポップ
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			elem := stack[len(stack)-1]
			elem.Children = append(elem.Children, CharData(tok))
		}
	}
	return result, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	node, err := NewTree(dec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", node)
}
