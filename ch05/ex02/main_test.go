package main

import (
	"testing"

	"strings"

	"reflect"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{name: "simple html", input: "<html><head></head><body></body></html>", expected: map[string]int{"html": 1, "head": 1, "body": 1}},
		{name: "simple html with newline",
			input: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			expected: map[string]int{"html": 1, "head": 1, "body": 1}},
		{name: "html with some objects",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="hogehoge.com">hogehoge</a>
		<div class="fuga">
			<p>fuga</p>
			<p>fugafuga</p>
		</div>
	</body>
</html>`,
			expected: map[string]int{"html": 1, "head": 1, "body": 1, "a": 1, "div": 1, "p": 2}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := make(map[string]int)
			outline(actual, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("outline %s expects %v but actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
