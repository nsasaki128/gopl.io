package main

import (
	"flag"
	"fmt"
	"time"
)

func pipeline(num int) {
	var out <-chan struct{}
	for i := 0; i < num; i++ {
		ch := make(chan struct{})
		go func(in chan<- struct{}, out <-chan struct{}) {
			in <- <-out
		}(ch, out)
		out = ch
	}

}

var num = flag.Int("num", 100000, "pipeline num")

func main() {
	flag.Parse()
	start := time.Now()
	fmt.Printf("pipeline num=%d\n", *num)
	//for local test
	//fmt.Printf("Create a pipeline.\n")
	pipeline(*num)

	//fmt.Printf("Send `struct{}{}` to the pipeline.\n")
	fmt.Printf("%s have passed\n", time.Since(start))

}
