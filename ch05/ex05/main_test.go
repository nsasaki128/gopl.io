package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	testCases := []struct {
		name                          string
		input                         string
		expectedWords, expectedImages int
	}{
		{name: "no link simple html", input: "<html><head></head><body></body></html>", expectedWords: 0, expectedImages: 0},
		{name: "1 word",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="hogehoge.com">hogehoge</a>
	</body>
</html>`,
			expectedWords: 1, expectedImages: 0},
		{name: "1 image",
			input: `
<html>
	<head>
	</head>
	<body>
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
	</body>
</html>`,
			expectedWords: 0, expectedImages: 1},
		{name: "some links",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="hogehoge.com">hoge hoge</a>
		<a href="/top">top</a>
		<a href="/link/books">books</a>
	</body>
</html>`,
			expectedWords: 4, expectedImages: 0},
		{name: "some image",
			input: `
<html>
	<head>
	</head>
	<body>
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
	</body>
</html>`,
			expectedWords: 0, expectedImages: 3},
		{name: "some links including link, script and, img",
			input: `
<html>
	<head>
		<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
	</head>
	<body>
		<a href="hogehoge.com">hoge hoge</a>
		<a href="/top">top</a>
		<a href="/link/books">books</a>
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
		<img class="browsers" src="../images/b_i4.gif" alt="Internet Explorer4">
	</body>
</html>`,
			expectedWords: 4, expectedImages: 2},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actualWords, actualImages := countWordsAndImages(doc)
			if actualWords != testCase.expectedWords {
				t.Errorf("%s expects words %d but actual %d", testCase.input, testCase.expectedWords, actualWords)
			}
			if actualImages != testCase.expectedImages {
				t.Errorf("%s expects images %d but actual %d", testCase.input, testCase.expectedImages, actualImages)
			}
		})
	}
}
