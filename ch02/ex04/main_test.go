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
func TestIteratePopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := IteratePopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}

func TestShiftPopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := ShiftPopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(1<<63 - 1))
		PopCount(uint64(0))
		PopCount(uint64(127))
	}
}
func BenchmarkIteratePopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IteratePopCount(uint64(1<<63 - 1))
		IteratePopCount(uint64(0))
		IteratePopCount(uint64(127))
	}
}
func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(uint64(1<<63 - 1))
		ShiftPopCount(uint64(0))
		ShiftPopCount(uint64(127))
	}
}


