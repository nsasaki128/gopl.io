package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

const (
	urlHeader = "http://"
)

func main() {
	fetchUrls(os.Stdout, os.Stderr, os.Args[1:])
}

func fetchUrls(outStream io.Writer, errStream io.Writer, urls []string) {
	for _, url := range urls {
		fetch(outStream, errStream, addUrlHeaderIfNeeded(url))
	}
}

func fetch(outStream io.Writer, errStream io.Writer, url string){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(errStream, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(outStream, resp.Body)
	resp.Body.Close()
	if err != nil{
		fmt.Fprintf(errStream, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

func addUrlHeaderIfNeeded(url string) string {
	if !strings.HasPrefix(url, urlHeader) {
		url = urlHeader + url
	}
	return url
}
