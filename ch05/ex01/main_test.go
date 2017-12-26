package main

import (
	"testing"

	"strings"

	"reflect"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "no link simple html", input: "<html><head></head><body></body></html>", expected: nil},
		{name: "no link html with newline",
			input: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			expected: nil},
		{name: "1 link",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="hogehoge.com">hogehoge</a>
	</body>
</html>`,
			expected: []string{"hogehoge.com"}},
		{name: "some links",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="hogehoge.com">hogehoge</a>
		<a href="/top">top</a>
		<a href="/link/books">books</a>
	</body>
</html>`,
			expected: []string{"hogehoge.com", "/top", "/link/books"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := visit(nil, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("links %s expects %v but actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
