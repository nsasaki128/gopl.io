// Echo0 は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(echo(os.Args))
}


func echo(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
