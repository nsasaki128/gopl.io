package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type result struct {
	url      string
	filename string
	size     int64
}

func main() {
	results := make(chan result, len(os.Args[1:]))
	done := make(chan struct{})
	for _, arg := range os.Args[1:] {
		go func(arg string) {
			filename, n, err := fetch(arg, done)
			if err != nil {
				fmt.Fprintf(os.Stderr, "title: %v\n", err)
				return
			}
			results <- result{url: arg, filename: filename, size: n}
		}(arg)
	}
	r := <-results
	close(done)
	fmt.Fprintf(os.Stdout, "url: %s\tfilename: %s\nsize: %d\n", r.url, r.filename, r.size)
}
func fetch(url string, done <-chan struct{}) (filename string, n int64, err error) {
	ctx, cancel := context.WithCancel(context.Background())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	cancel()
	return local, n, err

}
