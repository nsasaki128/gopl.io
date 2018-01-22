package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var c WordCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s. I'm John.", name)
	fmt.Println(c)
	var c2 LineCounter
	var name2 = "Dolly"
	fmt.Fprintf(&c2, "hello, %s.\nI'm John.\n", name2)
	fmt.Println(c2)
}
