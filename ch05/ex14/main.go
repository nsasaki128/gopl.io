package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	breadthFirst(workFileSystem, os.Args[1:])
}

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

func workFileSystem(dir string) []string {
	fmt.Println(dir)
	match, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Fatal(err)
	}

	var ret []string
	for _, m := range match {
		//last element
		if strings.HasPrefix(filepath.Base(m), ".") {
			continue
		}
		info, err := os.Stat(m)
		if err != nil {
			continue
		}
		if !info.IsDir() {
			continue
		}
		ret = append(ret, m)
	}
	return ret
}
