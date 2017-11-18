package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(echo(os.Args))
}
func echo(args []string) string {
	return strings.Join(args, " ")

}
