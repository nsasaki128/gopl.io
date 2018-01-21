package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	testCases := []struct {
		name      string
		inputHtml string
		inputName []string
		expected  string
	}{
		{name: "no match", inputHtml: "<html><head></head><div>test</div><body></body></html>", inputName: []string{"a"}, expected: ""},
		{name: "1 match for 1 input", inputHtml: "<html><head></head><div>test</div><body></body></html>", inputName: []string{"div"}, expected: "div"},
		{name: "2 match for 1 input", inputHtml: "<html><head></head><div>test1</div><div>test2</div><body></body></html>", inputName: []string{"div"}, expected: "divdiv"},
		{name: "1 match for 2 input", inputHtml: "<html><head></head><div>test</div><body></body></html>", inputName: []string{"div", "a"}, expected: "div"},
		{name: "2 match for 2 inputs", inputHtml: "<html><head></head><div>test</div><a href=\"hogepiyo\">fuga</a><body></body></html>", inputName: []string{"div", "a"}, expected: "divahrefhogepiyo"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input, _ := html.Parse(strings.NewReader(testCase.inputHtml))
			actualElement := ElementsByTagName(input, testCase.inputName...)
			var actual string
			for _, a := range actualElement {
				actual += a.Data
				for _, aa := range a.Attr {
					actual += aa.Key
					actual += aa.Val
				}
			}

			if actual != testCase.expected {
				t.Errorf("inputHtml %s expects %s but acutal is %s\n", testCase.inputHtml, testCase.expected, actual)

			}

		})
	}
}
