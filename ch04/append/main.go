package main

import "fmt"

func appendInt(x []int, y ...int) []int{
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// 拡大する余地がある。スライスを拡張する。
		z = x[:zlen]
	} else {
		//十分な容量がない。新たな配列を割り当てる。
		// 計算量を線形に均すために倍に拡張する。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):] , y)
	return z
}

func main()  {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}