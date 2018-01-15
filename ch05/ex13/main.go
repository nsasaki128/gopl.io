package main

import (
	"fmt"
	"log"
	"os"

	"net/url"

	"net/http"

	"io/ioutil"

	"path/filepath"

	"gopl.io/ch05/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	filtered := selectSameHost(url, list)
	for _, link := range filtered {
		save(link)
	}
	return filtered
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func selectSameHost(original string, list []string) []string {
	originalUrl, err := url.Parse(original)
	if err != nil {
		log.Println(err)
		return nil
	}

	var filtered []string

	for _, target := range list {
		compareUrl, err := url.Parse(target)
		if err != nil {
			log.Println(err)
			continue
		}
		if originalUrl.Host == compareUrl.Host {
			filtered = append(filtered, target)
		}
	}
	return filtered
}

func save(file string) {
	resp, err := http.Get(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	doc, err := url.Parse(file)
	if err != nil {
		log.Println(err)
		return
	}
	dir := filepath.Join("./", doc.Host, filepath.Clean(doc.Path))
	//for dividing dir path and file name
	base := "content"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(dir, base), bs, os.ModePerm)

	if err != nil {
		log.Println(err)
		return
	}
}
