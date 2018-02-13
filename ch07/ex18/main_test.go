package main

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Node
	}{
		{name: "1 element", input: "<html>hoge</html>", expected: &Element{
			Type:     xml.Name{Space: "", Local: "html"},
			Attr:     []xml.Attr{},
			Children: []Node{CharData("hoge")},
		},
		},
		{name: "2 elements", input: "<html><head>hoge</head><body>fuga</body></html>", expected: &Element{
			Type: xml.Name{Space: "", Local: "html"},
			Attr: []xml.Attr{},
			Children: []Node{
				&Element{
					Type:     xml.Name{Space: "", Local: "head"},
					Attr:     []xml.Attr{},
					Children: []Node{CharData("hoge")},
				},
				&Element{
					Type:     xml.Name{Space: "", Local: "body"},
					Attr:     []xml.Attr{},
					Children: []Node{CharData("fuga")},
				},
			},
		},
		},
		{name: "3 elements", input: "<html><head>hoge</head><body>fuga<a>piyo</a></body></html>", expected: &Element{
			Type: xml.Name{Space: "", Local: "html"},
			Attr: []xml.Attr{},
			Children: []Node{
				&Element{
					Type:     xml.Name{Space: "", Local: "head"},
					Attr:     []xml.Attr{},
					Children: []Node{CharData("hoge")},
				},
				&Element{
					Type: xml.Name{Space: "", Local: "body"},
					Attr: []xml.Attr{},
					Children: []Node{
						CharData("fuga"),
						&Element{
							Type:     xml.Name{Space: "", Local: "a"},
							Attr:     []xml.Attr{},
							Children: []Node{CharData("piyo")},
						},
					},
				},
			},
		},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dec := xml.NewDecoder(strings.NewReader(testCase.input))
			actual, err := NewTree(dec)
			if err != nil {
				t.Errorf("%s expects success but fails", testCase.input)
			} else if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("%s expects as %#v but actual is %#v", testCase.input, testCase.expected, actual)
			}
		})
	}

}
