package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var result []string

	for i, arg := range os.Args[1:] {
		result = append(result, strconv.Itoa(i+1) + " " + arg + "\n")
	}
	fmt.Print(strings.Join(result[0:], ""))
}
