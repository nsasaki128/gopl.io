package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "empty", input: "", expected: 0},
		{name: "1 word", input: "test", expected: 1},
		{name: "1 word with white", input: "test ", expected: 1},
		{name: "2 words", input: "test desu", expected: 2},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var c WordCounter
			c.Write([]byte(testCase.input))
			actual := int(c)
			if actual != testCase.expected {
				t.Errorf("<%s> expected %d but actual is %d\n", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func TestLineCounter(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "empty", input: "", expected: 0},
		{name: "1 word", input: "test", expected: 1},
		{name: "1 word with white", input: "test ", expected: 1},
		{name: "2 words", input: "test desu", expected: 1},
		{name: "1 line finish at newline", input: "test desu\n", expected: 1},
		{name: "2 lines", input: "test desu\nI'm happy", expected: 2},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var c LineCounter
			c.Write([]byte(testCase.input))
			actual := int(c)
			if actual != testCase.expected {
				t.Errorf("<%s> expected %d but actual is %d\n", testCase.input, testCase.expected, actual)
			}
		})
	}
}
