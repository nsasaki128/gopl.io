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
func TestClearPopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := ClearPopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}
func TestDivideAndConquerPopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := DivideAndConquerPopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}

var output int

func benchmark_PopCount(b *testing.B, v uint64) {
	var a int
	for i := 0; i < b.N; i++ {
		a += PopCount(v)
	}
	output = a
}
func benchmark_IteratePopCount(b *testing.B, v uint64) {
	var a int
	for i := 0; i < b.N; i++ {
		a += IteratePopCount(v)
	}
	output = a
}
func benchmark_ShiftPopCount(b *testing.B, v uint64) {
	var a int
	for i := 0; i < b.N; i++ {
		a += ShiftPopCount(v)
	}
	output = a
}
func benchmark_ClearPopCount(b *testing.B, v uint64) {
	var a int
	for i := 0; i < b.N; i++ {
		a += ClearPopCount(v)
	}
	output = a
}
func benchmark_DivideAndConquerPopCount(b *testing.B, v uint64) {
	var a int
	for i := 0; i < b.N; i++ {
		a += DivideAndConquerPopCount(v)
	}
	output = a
}

const all1 = uint64(1<<63 - 1)

func BenchmarkPopCount_0(b *testing.B)        { benchmark_PopCount(b, uint64(0)) }
func BenchmarkIteratePopCount_0(b *testing.B) { benchmark_IteratePopCount(b, uint64(0)) }
func BenchmarkShiftPopCount_0(b *testing.B)   { benchmark_ShiftPopCount(b, uint64(0)) }
func BenchmarkClearPopCount_0(b *testing.B)   { benchmark_ClearPopCount(b, uint64(0)) }
func BenchmarkDivideAndConquerPopCount_0(b *testing.B) {
	benchmark_DivideAndConquerPopCount(b, uint64(0))
}
func BenchmarkPopCount_some1(b *testing.B)        { benchmark_PopCount(b, uint64(127)) }
func BenchmarkIteratePopCount_some1(b *testing.B) { benchmark_IteratePopCount(b, uint64(127)) }
func BenchmarkShiftPopCount_some1(b *testing.B)   { benchmark_ShiftPopCount(b, uint64(127)) }
func BenchmarkClearPopCount_some1(b *testing.B)   { benchmark_ClearPopCount(b, uint64(127)) }
func BenchmarkDivideAndConquerPopCount_some1(b *testing.B) {
	benchmark_DivideAndConquerPopCount(b, uint64(127))
}
func BenchmarkPopCount_all1(b *testing.B)                 { benchmark_PopCount(b, all1) }
func BenchmarkIteratePopCount_all1(b *testing.B)          { benchmark_IteratePopCount(b, all1) }
func BenchmarkShiftPopCount_all1(b *testing.B)            { benchmark_ShiftPopCount(b, all1) }
func BenchmarkClearPopCount_all1(b *testing.B)            { benchmark_ClearPopCount(b, all1) }
func BenchmarkDivideAndConquerPopCount_all1(b *testing.B) { benchmark_DivideAndConquerPopCount(b, all1) }
