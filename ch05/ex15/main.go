package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	vals := []int{}
	for _, v := range os.Args[1:] {
		if i, err := strconv.Atoi(v); err == nil {
			vals = append(vals, i)
		}
	}
	max, err := max(vals...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("max(%v) = %v\n", vals, max)

	min, err := min(vals...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("min(%v) = %v\n", vals, min)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("length of func max is more than 1")
	}
	max := math.MinInt64
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("length of func min is more than 1")
	}
	min := math.MaxInt64
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min, nil
}

func max2(m int, vals ...int) int {
	max := m
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min2(m int, vals ...int) int {
	min := m
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
