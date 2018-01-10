package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expand(string(bs), func(s string) string { return s + ":\n" }))
}

var r = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	bs := r.ReplaceAllFunc([]byte(s), func(bs []byte) []byte {
		return []byte(f(string(bs[1:])))
	})

	return string(bs)
}
