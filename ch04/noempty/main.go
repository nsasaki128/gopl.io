package main

import "fmt"

// noempty は空文字でない文字列を保持するスライスを返します。
// 基底配列は呼び出し中に修正されます。
func noempty(strings []string)  []string{
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0] // もとの長さ0のスライス
	for _, s := range strings{
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main()  {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)
}

