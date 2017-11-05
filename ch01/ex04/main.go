package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
	"io"
)

var out io.Writer = os.Stdout

func main() {
	files := os.Args[1:]
	writeDupLineAndFiles(files)
}

func writeDupLineAndFiles(files []string) {
	counts := make(map[string]map[string]int)
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

	writeCounts(counts)
}

func countLines(f *os.File, fileName string, counts map[string]map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		countLine(input.Text(), fileName, counts)
	}
	//注意: input.Err()からのエラーの可能性を無視している
}

func countLine(text string,fileName string, counts map[string]map[string]int) {
	if counts[text]== nil {
		 counts[text] = make(map[string]int)
	}
	counts[text][fileName]++
}

func writeCounts(counts map[string]map[string]int) {
	//for testing results, here I sorted keys
	var lines []string
	for line := range counts {
		lines = append(lines, line)
	}
	sort.Strings(lines)

	for _, line := range lines {
		var sum = 0
		var fileNames []string
		for fileName, n := range counts[line] {
			sum += n
			fileNames = append(fileNames, fileName)
		}
		sort.Strings(fileNames)

		if sum > 1 {
			fmt.Fprintf(out, "%d\t%s\t%s\n", sum, line, strings.Join(fileNames, "\t"))
		}
	}
}