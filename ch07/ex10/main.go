package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type runeSort []rune

func (s runeSort) Len() int           { return len(s) }
func (s runeSort) Less(i, j int) bool { return s[i] < s[j] }
func (s runeSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main plindromes")
	}
	for _, arg := range os.Args[1:] {
		if IsPalindrome(runeSort([]rune(arg))) {
			fmt.Printf("%q is palindrome!\n", arg)
		} else {
			fmt.Printf("%q is not palindrome.\n", arg)
		}
	}
}
