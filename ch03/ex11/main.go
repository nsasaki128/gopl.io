package main

import (
"bufio"
"fmt"
"os"
	"bytes"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

// comma は負でない10進表記整数文字列にカンマを挿入します。
func comma(s string) string {
	if len(s) < 3 {
		return s
	}
	sp := strings.Split(s, ".")
	integer := sp[0]
	fractional := ""
	if len(sp) > 1 {
		fractional = "." + sp[1]
	}

	var buf bytes.Buffer
	if integer[0] == '-' || integer[0] == '+' {
		buf.WriteByte(integer[0])
		integer = integer[1:]
	}

	buf.WriteByte(integer[0])
	for i:=1; i < len(integer); i++ {
		if (len(integer)-i)%3 == 0{
			buf.WriteString(",")
		}
		buf.WriteByte(integer[i])
	}
	buf.WriteString(fractional)
	return buf.String()
}

