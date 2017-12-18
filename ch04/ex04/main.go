package main

import "fmt"

func main()  {
	input := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Printf("original:\t%v\n", input)
	for i := 1; i < len(input); i++ {
		fmt.Printf("rotate %d:\t%v\n", i, rotate(input, i))
	}
}

func gcd (x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func rotate(input []int, rotate int) []int {
	output := make([]int, len(input))
	copy(output, input)
	n := gcd(len(input), rotate)
	for i := 0; i < n; i++ {
		//rotateEach
		prev := output[i]
		for j := 0; j < len(input)/n; j++ {
			next := (i+(j+1)*rotate)%len(input)
			tmp := output[next]
			output[next] = prev
			prev = tmp
		}
	}
	return output
}


