package main

import "fmt"

func main() {
	fmt.Println(inc(10))
}

func inc(i int) (res int) {
	defer func() {
		recover()
		res = i + 1
	}()
	panic(nil)
}
