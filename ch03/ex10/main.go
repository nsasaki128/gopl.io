package main

import (
"bufio"
"fmt"
"os"
	"bytes"
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
	var buf bytes.Buffer
	buf.WriteByte(s[0])
	for i:=1; i < len(s); i++ {
		if (len(s)-i)%3 == 0{
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

