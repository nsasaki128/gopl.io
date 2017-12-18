package main

import "fmt"

func main(){
	test := []int{1, 2, 3, 4, 5}
	fmt.Printf("original: ")
	fmt.Println(test)
	reverse(test)
	fmt.Printf("reverse: ")
	fmt.Println(test)

}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
