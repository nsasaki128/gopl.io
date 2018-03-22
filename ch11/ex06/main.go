package main

// pc[i]はiのポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント（1が設定されているビット数）を返します。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func IteratePopCount(x uint64) int {
	var out int
	for i := uint(0); i < 8; i++ {
		out += int(pc[byte(x>>(i*8))])
	}
	return out
}

func ShiftPopCount(x uint64) int {
	var out int
	for i := uint(0); i < 64; i++ {
		out += int(x>>i) & 1
	}
	return out
}

func ClearPopCount(x uint64) int {
	var out int
	for x != 0 {
		x &= (x - 1)
		out++
	}
	return out
}

func DivideAndConquerPopCount(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x >> 4) & 0x0F0F0F0F0F0F0F0F)
	x = (x & 0x00FF00FF00FF00FF) + ((x >> 8) & 0x00FF00FF00FF00FF)
	x = (x & 0x0000FFFF0000FFFF) + ((x >> 16) & 0x0000FFFF0000FFFF)
	x = (x & 0x00000000FFFFFFFF) + ((x >> 32) & 0x00000000FFFFFFFF)
	return int(x)
}
