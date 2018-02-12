package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"strconv"

	"gopl.io/ch07/ex15/eval"
)

func main() {
	fmt.Println("input Expr")
	r := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		fmt.Print("> ")
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatal(err)
		}
		expr, err := eval.Parse(string(line))
		if err != nil {
			fmt.Printf("Invalid expression: %s\n", line)
			continue
		}

		vars := map[eval.Var]bool{}
		err = expr.Check(vars)
		if err != nil {
			fmt.Printf("Invalid expression: %s\n", err)
			continue
		}

		env := eval.Env{}

		for v := range vars {
			for {
				fmt.Printf("%s = ", v)
				line, _, err := r.ReadLine()
				if err == io.EOF {
					return
				} else if err != nil {
					log.Fatal(err)
				}

				f, err := strconv.ParseFloat(string(line), 64)
				if err != nil {
					fmt.Printf("Invalid value: %s\n", err)
					continue
				}
				env[v] = f
				break
			}
		}
		fmt.Printf("%f\n", expr.Eval(env))
	}
}
