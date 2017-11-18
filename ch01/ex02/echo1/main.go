// Echo1 は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
		fmt.Print(echoWithIndex(os.Args[1:]))
	}

func echoWithIndex(args []string) string {
	var s, sep, nl string
	sep = " "
	nl = "\n"
	for i := 0; i < len(args); i++ {
		s += strconv.Itoa(i+1) + sep + args[i] + nl
	}
	return s
}
