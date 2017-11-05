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

	for _, url := range os.Args[1:] {
		writeRespBody(url)
	}
}

func writeRespBody(url string) {

	url = addUrlHeaderIfNeeded(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "HTTP STATUS: %s\n", resp.Status)

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

}

func addUrlHeaderIfNeeded(url string) string {
	if !strings.HasPrefix(url, urlHeader) {
		url = urlHeader + url
	}
	return url
}
