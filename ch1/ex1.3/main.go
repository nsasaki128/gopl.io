package main

import (
	"os"
	"fmt"
	"strings"
)

func echo1(argv []string){
	s, sep := "", ""
	for i := 1; i < len(argv); i++ {
		s += sep + argv[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(argv []string){
	s, sep := "", ""
	for _, arg := range argv[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(argv []string){
	fmt.Println(strings.Join(argv[1:], " "))
}


func  main()  {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}
