package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, arg := range os.Args[1:] {
		filename, n, err := fetch(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "filename: %s\nsize: %d\n", filename, n)
	}
}
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
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
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err

}
