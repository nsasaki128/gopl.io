package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}


	for line, names := range counts {
		var sum = 0
		var fileNames = ""
		for fileName, n := range names {
			sum += n
			fileNames += "\t" + fileName
		}

		if sum > 1 {
			fmt.Printf("%d\t%s%s\n", sum, line, fileNames)
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if(counts[input.Text()] == nil) {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][fileName]++
	}
	//注意: input.Err()からのエラーの可能性を無視している
}
