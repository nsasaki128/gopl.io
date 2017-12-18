package main

import (
	"unicode/utf8"
	"unicode"
	"fmt"
)

func main()  {
	input1 := "こんにちは　　　世界"
	input2 := "こんにちは   世界"
	input3 := "こんにちは　 　世界"
	fmt.Printf("%s\n", string(compress([]byte(input1))))
	fmt.Printf("%s\n", string(compress([]byte(input2))))
	fmt.Printf("%s\n", string(compress([]byte(input3))))
}

func compress(bytes []byte) []byte {
	out := bytes[:0]
	i := 0
	var prev rune
	for i < len(bytes) {
		r, d := utf8.DecodeRune(bytes[i:])
		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+d]...)
		} else if !unicode.IsSpace(prev) {
			out = append(out, []byte(" ")...)
		}
		prev = r
		i += d
	}
	return out
}
