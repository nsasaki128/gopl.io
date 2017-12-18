package main

import (
	"unicode/utf8"
	"fmt"
)

func main()  {
	input1 := []byte("こんにちは 世界")
	input2 := []byte("hello, 世界")
	fmt.Printf("%s\n", string(input1))
	fmt.Printf("%s\n", string(input2))
	reverseByte(input1)
	reverseByte(input2)
	fmt.Printf("%s\n", string(input1))
	fmt.Printf("%s\n", string(input2))
}

func reverseByte(bytes []byte) {
	if len(bytes)  <= 0 {
		return
	}
	_, d := utf8.DecodeRune(bytes)
	reverse(bytes[:d])
	reverse(bytes[d:])
	reverse(bytes)
	reverseByte(bytes[:len(bytes)-d])

}


func reverse(bytes []byte) {
	for i, j := 0, len(bytes) - 1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
