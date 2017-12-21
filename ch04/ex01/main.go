package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s, s1, s2", os.Args[0])
		os.Exit(1)
	}
	dif := difBitCountInSha256([]byte(os.Args[1]), []byte(os.Args[2]))
	fmt.Fprintf(os.Stdout, "sha256 hash bit count differece between %s and %s is %d\n", os.Args[1], os.Args[2], dif)
}

func DivideAndConquerPopCount(x uint8) int {
	x = (x & 0x55) + ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	x = (x & 0x0F) + ((x >> 4) & 0x0F)
	return int(x)
}

func difBitCountInSha256(b1, b2 []byte) int {
	c1 := sha256.Sum256([]byte(b1))
	c2 := sha256.Sum256([]byte(b2))
	var sum int
	for i := 0; i < len(c1); i++ {
		sum += DivideAndConquerPopCount(c1[i] ^ c2[i])
	}
	return sum
}
