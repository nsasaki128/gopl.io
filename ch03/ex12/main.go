package main

import (
	"os"
	"fmt"
)

func main()  {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s, s1, s2", os.Args[0])
		os.Exit(1)
	}
	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Fprintf(os.Stdout, "%s and %s are anagram\n", os.Args[1], os.Args[2])
	} else {
		fmt.Fprintf(os.Stdout, "%s and %s are not anagram\n", os.Args[1], os.Args[2])
	}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	c := map[rune] int{}
	for _, r := range s1 {
		c[r]++
	}
	for _, r := range s2 {
		c[r]--
		if c[r] < 0 {
			return false
		}
	}

	return true
}
