package main

import "sync"

// pc[i]はiのポピュレーションカウントです。
var pc [256]byte
var popCountInitOnce sync.Once

func popCountInit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント（1が設定されているビット数）を返します。
func PopCount(x uint64) int {
	popCountInitOnce.Do(popCountInit)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
