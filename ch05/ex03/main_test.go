package main

import (
	"testing"

	"strings"

	"reflect"

	"golang.org/x/net/html"
)

func TestFindtexts(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "only title", input: "<html><head><title>this is title</title></head><body></body></html>", expected: []string{"this is title"}},
		{name: "only body", input: "<html><head></head><body><h1>main body</h1></body></html>", expected: []string{"main body"}},
		{name: "style and script",
			input: "<html>" +
				"<head>" +
				"<style type=\"text/css\"><!--body {font-size: 14px;}--></style>" +
				"</head>" +
				"<body>" +
				"<script>alert(\"This is alert!!\")</script>" +
				"</body>" +
				"</html>",
			expected: nil},
		{name: "html with some objects",
			input: "<html>" +
				"<head><title>this is title</title>" +
				"<style type=\"text/css\"><!--body {font-size: 14px;}--></style>" +
				"</head>" +
				"<body>" +
				"<script>alert(\"This is alert!!\")</script>" +
				"<body>これはボディ<body>" +
				"</body>" +
				"</html>",
			expected: []string{"this is title", "これはボディ"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := findtexts(nil, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("outline %s expects %v but actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
