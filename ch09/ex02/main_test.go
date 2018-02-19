package main

import (
	"testing"
)

var testCases = []struct {
	name     string
	input    uint64
	expected int
}{
	{name: "input 0", input: 0, expected: 0},
	{name: "input 1", input: 1, expected: 1},
	{name: "input 2", input: 2, expected: 1},
	{name: "input 3", input: 3, expected: 2},
	{name: "input 4", input: 4, expected: 1},
	{name: "input MAX", input: 1<<63 - 1, expected: 63},
}

func TestPopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := PopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}
