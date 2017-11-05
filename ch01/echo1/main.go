// Echo1 は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
