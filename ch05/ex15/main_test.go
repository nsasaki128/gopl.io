package main

import (
	"math"
	"testing"
)

var (
	testCases = []struct {
		name        string
		input       []int
		isAnsExist  bool
		expectedMax int
		expectedMin int
	}{
		{name: "no input", input: []int{}, isAnsExist: false, expectedMax: 0, expectedMin: 0},
		{name: "1 input", input: []int{1}, isAnsExist: true, expectedMax: 1, expectedMin: 1},
		{name: "2 input", input: []int{1, 100}, isAnsExist: true, expectedMax: 100, expectedMin: 1},
		{name: "3 input with minus", input: []int{1, 100, -100}, isAnsExist: true, expectedMax: 100, expectedMin: -100},
		{name: "input with Maxvalue", input: []int{1, math.MaxInt64, -100}, isAnsExist: true, expectedMax: math.MaxInt64, expectedMin: -100},
		{name: "input with Minvalue", input: []int{1, 100, math.MinInt64}, isAnsExist: true, expectedMax: 100, expectedMin: math.MinInt64},
	}
)

func TestMax(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := max(testCase.input...)
			if (err != nil) != !testCase.isAnsExist {
				t.Errorf("input %v expects answer %v but actual %v", testCase.input, testCase.isAnsExist, err != nil)
			}
			if actual != testCase.expectedMax {
				t.Errorf("input %v expects max %d but actual %d", testCase.input, testCase.expectedMax, actual)
			}
		})
	}
}

func TestMin(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := min(testCase.input...)
			if (err != nil) != !testCase.isAnsExist {
				t.Errorf("input %v expects answer %v but actual %v", testCase.input, testCase.isAnsExist, err != nil)
			}
			if actual != testCase.expectedMin {
				t.Errorf("input %v expects min %d but actual %d", testCase.input, testCase.expectedMin, actual)
			}
		})
	}
}

func TestMax2(t *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.input) <= 0 {
			continue
		}
		t.Run(testCase.name, func(t *testing.T) {
			actual := max2(testCase.input[0], testCase.input...)
			if actual != testCase.expectedMax {
				t.Errorf("input %v expects max %d but actual %d", testCase.input, testCase.expectedMax, actual)
			}
		})
	}
}

func TestMin2(t *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.input) <= 0 {
			continue
		}
		t.Run(testCase.name, func(t *testing.T) {
			actual := min2(testCase.input[0], testCase.input...)
			if actual != testCase.expectedMin {
				t.Errorf("input %v expects min %d but actual %d", testCase.input, testCase.expectedMin, actual)
			}
		})
	}
}
