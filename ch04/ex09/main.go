package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
)

func main() {
	f := wordfreq(os.Stdin)

	fmt.Println("word\tcount")
	for k, v:= range f {
		fmt.Printf("%s\t%d\n", k, v)
	}

}

func wordfreq(r io.Reader) map[string]int {
	result := make(map[string]int)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		result[s.Text()]++
	}
	return result
}