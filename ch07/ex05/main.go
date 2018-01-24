package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type limitReader struct {
	r io.Reader
	n int64
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[:r.n]
	}
	n, err = r.r.Read(p)
	r.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	reader := bytes.NewReader([]byte("Hello world"))

	data, err := ioutil.ReadAll(LimitReader(reader, 2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	data, err = ioutil.ReadAll(LimitReader(reader, 2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
