package main

import (
	"gopl.io/ch02/ex02/tempconv"
	"os"
	"io"
	"bufio"
	"strconv"
	"fmt"
	"gopl.io/ch02/ex02/lengthconv"
	"gopl.io/ch02/ex02/weightconv"
)


func main () {
	if len(os.Args) <= 1 {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan(){
			printConvValues(in.Text(), os.Stdout)
		}

	}else{
		for _, arg := range os.Args[1:] {
			printConvValues(arg , os.Stdout)
		}
	}
}

func printConvValues(in string, out io.Writer)  {
	t, err := strconv.ParseFloat(in, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Fprintf(out, "%s = %s\n%s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	ft := lengthconv.Feet(t)
	m  := lengthconv.Meter(t)
	fmt.Fprintf(out, "%s = %s\n%s = %s\n", ft, lengthconv.FToM(ft), m, lengthconv.MToF(m))

	lb := weightconv.Pound(t)
	kg := weightconv.Kilogram(t)
	fmt.Fprintf(out, "%s = %s\n%s = %s\n", lb, weightconv.PToK(lb), kg, weightconv.KToP(kg) )

	fmt.Fprintln(out)

}