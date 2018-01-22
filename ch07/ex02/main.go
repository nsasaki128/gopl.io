package main

import (
	"bytes"
	"fmt"
	"io"
)

type wrapper struct {
	counter int64
	w       io.Writer
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var result wrapper
	result.w = w
	return &result, &(result.counter)
}

func (c *wrapper) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.counter += int64(n)
	return n, err
}

func main() {
	buf := bytes.Buffer{}
	cw, counter := CountingWriter(&buf)
	bs := []byte("Hello, world!")
	cw.Write(bs)
	fmt.Println(cw)
	fmt.Println(*counter)

}
