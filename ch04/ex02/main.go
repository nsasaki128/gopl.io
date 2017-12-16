package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"flag"
	"crypto/sha512"
	"bufio"
)

var hashType = flag.String("t", "sha256", "hash function for (sha256, sha384, sha512)")
func main()  {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		switch *hashType {
		case "sha256":
			fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256(input.Bytes()))
		case "sha384":
			fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum384(input.Bytes()))
		case "sha512":
			fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512(input.Bytes()))
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}


