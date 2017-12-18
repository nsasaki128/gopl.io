package main

import "fmt"

func main()  {
	input1 := []string{"hoge", "fuga", "piyo"}
	input2 := []string{"hoge", "hoge", "fuga", "piyo", "fuga"}
	fmt.Printf("original:\t%v\n", input1)
	fmt.Printf("delete succeed duplicated:\t%v\n", deleteSucDup(input1))
	fmt.Printf("original:\t%v\n", input2)
	fmt.Printf("delete succeed duplicated:\t%v\n", deleteSucDup(input2))
}

func deleteSucDup(strings []string) []string {
	out := strings[:0]
	out = append(out, strings[0])
	for _, s := range strings {
		if s != out[len(out)-1]{
			out = append(out, s)
		}
	}
	return out
}


