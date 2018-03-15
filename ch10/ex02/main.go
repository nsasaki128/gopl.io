package main

import (
	"fmt"
	"io"
	"log"

	"os"

	"gopl.io/ch10/ex02/archive"
	_ "gopl.io/ch10/ex02/archive/tar"
	_ "gopl.io/ch10/ex02/archive/zip"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	r, _, err := archive.Read(file)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintf(os.Stderr, "Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, r); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
