package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}
