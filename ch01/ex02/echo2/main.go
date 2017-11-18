// Echo2は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func  main()  {
	fmt.Print(echoWithIndex(os.Args[1:]))
}

func echoWithIndex(args []string) string {
	s, sep, nl := "", " ", "\n"
	for i, arg := range args {
		s += strconv.Itoa(i+1) + sep + arg + nl
	}
	return s
}
