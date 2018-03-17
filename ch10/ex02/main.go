package main

import (
	"fmt"
	"io"
	"log"

	"os"

	"bytes"
	"io/ioutil"
	"path/filepath"

	"gopl.io/ch10/ex02/archive"
	_ "gopl.io/ch10/ex02/archive/tar"
	_ "gopl.io/ch10/ex02/archive/zip"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage: need to add archived file and output directory.\n")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r, err := archive.Read(file)
	if err != nil {
		log.Fatalln(err)
	}
	output := os.Args[2]
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintf(os.Stderr, "Contents of %s:\n", hdr.Name)
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, r); err != nil {
			log.Fatal(err)
		}

		//Create file path
		p := filepath.Join(output, hdr.Name)
		d, _ := filepath.Split(p)
		if _, err = os.Stat(d); err != nil {
			os.MkdirAll(d, os.ModePerm)
		}

		//Write file
		if hdr.FileInfo.IsDir() {
			continue
		}
		if err = ioutil.WriteFile(p, buf.Bytes(), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
