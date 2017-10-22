package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func echo1(argv []string){
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(argv); i++ {
		s += sep + argv[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("echo1 %.2fs elapsed\n", time.Since(start).Seconds())
}

func echo2(argv []string){
	start := time.Now()
	s, sep := "", ""
	for _, arg := range argv[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("echo2 %.2fs elapsed\n", time.Since(start).Seconds())
}

func echo3(argv []string){
	start := time.Now()
	fmt.Println(strings.Join(argv[1:], " "))
	fmt.Printf("echo3 %.2fs elapsed\n", time.Since(start).Seconds())
	}


func  main()  {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)

}
