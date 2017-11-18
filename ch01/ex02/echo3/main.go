package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Print(echoWithIndex(os.Args[1:]))
}
func echoWithIndex(args []string) string {
	var results []string

	for i, arg := range args {
		results = append(results, strconv.Itoa(i+1) + " " + arg + "\n")
	}

	return strings.Join(results, "")
}