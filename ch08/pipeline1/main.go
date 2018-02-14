package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//Square
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//Printer （メインゴルーチン）
	for {
		fmt.Println(<-squares)
	}
}
