package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

func main (){
	counts := make(map[rune]int) //Unicode文字の数
	var utflen [utf8.UTFMax + 1]int // UTF-8エンコーディングの長さの数
	invalid := 0 //不正なUTF-8文字の数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //rune, nbyte, errorを返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTG-8 characters\n", invalid)
	}
}
