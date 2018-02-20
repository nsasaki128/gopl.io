package main

import (
	"fmt"
	"time"
)

func pingpong(in <-chan struct{}, out chan<- struct{}, done <-chan struct{}, result chan<- int) {
	for i := 0; ; i++ {
		select {
		case ping := <-in:
			select {
			case out <- ping:
			case <-done:
				result <- i
				return
			}
		case <-done:
			result <- i
			return
		}
	}
}

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	done := make(chan struct{})
	result := make(chan int)
	go func() {
		<-time.After(1 * time.Second)
		close(done)
	}()

	go pingpong(c1, c2, done, result)
	go pingpong(c2, c1, done, result)

	c1 <- struct{}{}
	r := <-result
	fmt.Printf("total %d pingpong\n", r)
}
