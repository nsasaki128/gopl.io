package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("original:\t%v\n", input)
	for i := 1; i < len(input); i++ {
		rotate(input, i)
		fmt.Printf("rotate %d:\t%v\n", i, input)
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func rotate(input []int, rotate int) {
	n := gcd(len(input), rotate)
	for i := 0; i < n; i++ {
		//rotateEach
		prev := input[i]
		for j := 0; j < len(input)/n; j++ {
			next := (i + (j+1)*rotate) % len(input)
			tmp := input[next]
			input[next] = prev
			prev = tmp
		}
	}
}
