// Echo2は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func  main()  {
	s, sep := "", " "
	for i, arg := range os.Args[1:] {
		s = strconv.Itoa(i+1) + sep + arg
		fmt.Println(s)
	}
}
