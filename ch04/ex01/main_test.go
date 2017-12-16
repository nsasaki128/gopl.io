package main

import (
	"testing"
)

func TestDivideAndConquerPopCount(t *testing.T) {
	var testCases = []struct {
		name     string
		input    uint8
		expected int
	}{
		{name: "input 0", input: 0, expected: 0},
		{name: "input 1", input: 1, expected: 1},
		{name: "input 2", input: 2, expected: 1},
		{name: "input 3", input: 3, expected: 2},
		{name: "input 4", input: 4, expected: 1},
		{name: "input 8", input: 8, expected: 1},
		{name: "input 9", input: 9, expected: 2},
		{name: "input 15", input: 15, expected: 4},
		{name: "input MAX", input: 1<<7 - 1, expected: 7},
	}
	for _, testCase := range testCases {
		result := DivideAndConquerPopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("case %s expected %d actual %d", testCase.name, testCase.expected, result)
		}
	}
}

func TestDifBitCountInSha256(t *testing.T) {
	var testCases = []struct {
		name     string
		input1   []byte
		input2   []byte
		expected bool
	}{
		{name: "same pattern0: nil", input1: nil, input2: nil, expected: true},
		{name: "same pattern1: one char", input1: []byte("x"), input2: []byte("x"), expected: true},
		{name: "same pattern2: some chars", input1: []byte("hogehoge"), input2: []byte("hogehoge"), expected: true},
		{name: "different pattern1", input1: []byte("X"), input2: []byte("x"), expected: false},
		{name: "different pattern2", input1: []byte("hogehoge"), input2: []byte("piyopiyo"), expected: false},
	}
	for _, testCase := range testCases {
		result := 0 == difBitCountInSha256(testCase.input1, testCase.input2)
		if result != testCase.expected {
			t.Errorf("case %s expected %t actual %t", testCase.name, testCase.expected, result)
		}
	}
}

