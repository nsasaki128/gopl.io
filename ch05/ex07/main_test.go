package main

import (
	"testing"

	"strings"

	"bytes"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "no link simple html", input: "<html><head></head><body></body></html>",
			expected: `<html>
  <head/>
  <body/>
</html>
`},
		{name: "1 href body",
			input: `
<html><head></head><body><a href="hogehoge.com">hogehoge</a></body></html>`,
			expected: `<html>
  <head/>
  <body>
    <a href="hogehoge.com">hogehoge</a>
  </body>
</html>
`},
		{name: "1 img",
			input: `<html><head></head><body><img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"></body></html>`,
			expected: `<html>
  <head/>
  <body>
    <img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"/>
  </body>
</html>
`},
		{name: "some links",
			input: `
<html><head></head><body><a href="hogehoge.com">hoge hoge</a><a href="/top">top</a><a href="/link/books">books</a></body></html>`,
			expected: `<html>
  <head/>
  <body>
    <a href="hogehoge.com">hoge hoge</a>
    <a href="/top">top</a>
    <a href="/link/books">books</a>
  </body>
</html>
`},
		{name: "some links including link, script and, img",
			input: `<html><head><link type="text/css" rel="stylesheet" href="/lib/godoc/style.css"><script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script></head><body><a href="hogehoge.com">hoge hoge</a><a href="/top">top</a><a href="/link/books">books</a><img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"><img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"></body></html>`,
			expected: `<html>
  <head>
    <link type="text/css" rel="stylesheet" href="/lib/godoc/style.css"/>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"/>
  </head>
  <body>
    <a href="hogehoge.com">hoge hoge</a>
    <a href="/top">top</a>
    <a href="/link/books">books</a>
    <img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"/>
    <img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4"/>
  </body>
</html>
`},
	}
	for _, testCase := range testCases {
		actual := new(bytes.Buffer)
		writer = actual
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			forEachNode(doc, startElement, endElement)

			if actual.String() != testCase.expected {
				t.Errorf("%s expects\n%s but actual\n%s", testCase.input, testCase.expected, actual.String())
			}
		})
	}
}
