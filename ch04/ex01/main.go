package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)


func main()  {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s, s1, s2", os.Args[0])
	}
	dif := difBitCountInSha256([]byte(os.Args[1]), []byte(os.Args[2]))
	fmt.Fprintf(os.Stdout, "sha256 hash bit count differece between %s and %s is %d\n", os.Args[1], os.Args[2], dif)
}

func DivideAndConquerPopCount(x uint64) int {
	x = (x&0x5555555555555555) + ((x>>1)&0x5555555555555555)
	x = (x&0x3333333333333333) + ((x>>2)&0x3333333333333333)
	x = (x&0x0F0F0F0F0F0F0F0F) + ((x>>4)&0x0F0F0F0F0F0F0F0F)
	x = (x&0x00FF00FF00FF00FF) + ((x>>8)&0x00FF00FF00FF00FF)
	x = (x&0x0000FFFF0000FFFF) + ((x>>16)&0x0000FFFF0000FFFF)
	x = (x&0x00000000FFFFFFFF) + ((x>>32)&0x00000000FFFFFFFF)
	return int(x)
}

func difBitCountInSha256(b1, b2 []byte) int {
	c1 := sha256.Sum256([]byte(b1))
	c2 := sha256.Sum256([]byte(b2))
	var sum int
	for i := 0; i < len(c1); i++ {
		sum += DivideAndConquerPopCount(uint64(c1[i])) ^ DivideAndConquerPopCount(uint64(c2[i]))
	}
	return sum
}
