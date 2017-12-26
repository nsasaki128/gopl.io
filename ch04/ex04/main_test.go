package main

import (
	"strconv"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		rotate   int
		expected []int
	}{
		{name: "12 rotates 0", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 0, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{name: "12 rotates 1", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 1, expected: []int{12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{name: "12 rotates 2", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 2, expected: []int{11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "12 rotates 3", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 3, expected: []int{10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "12 rotates 4", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 4, expected: []int{9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8}},
		{name: "12 rotates 5", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 5, expected: []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7}},
		{name: "12 rotates 6", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 6, expected: []int{7, 8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6}},
		{name: "12 rotates 7", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 7, expected: []int{6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4, 5}},
		{name: "12 rotates 8", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 8, expected: []int{5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4}},
		{name: "12 rotates 9", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 9, expected: []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3}},
		{name: "12 rotates 10", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 10, expected: []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2}},
		{name: "12 rotates 11", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, rotate: 11, expected: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1}},
	}

	for _, testCase := range testCases {
		rotate(testCase.input, testCase.rotate)
		t.Run(testCase.name, func(t *testing.T) {
			for i := 0; i < len(testCase.input); i++ {
				if testCase.input[i] != testCase.expected[i] {
					t.Errorf("error in case %s \nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected, testCase.input)
				}
			}
		})
	}
}

func TestGcd(t *testing.T) {

	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 1, y: 2, expected: 1},
		{x: 2, y: 1, expected: 1},
		{x: 2, y: 3, expected: 1},
		{x: 2, y: 4, expected: 2},
		{x: 12, y: 8, expected: 4},
	}
	for _, testCase := range testCases {
		t.Run(strconv.Itoa(testCase.x)+"and"+strconv.Itoa(testCase.y), func(t *testing.T) {
			actual := gcd(testCase.x, testCase.y)
			if actual != testCase.expected {
				t.Errorf("error in case %d and %d \nexpected:\t%d\nactual:\t%d\n", testCase.x, testCase.y, testCase.expected, actual)
			}
		})
	}
}
