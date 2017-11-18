// Echo2は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
)

func  main()  {
	fmt.Println(echo(os.Args))
}
func echo(args []string) string {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}
