package main

import (
	"testing"
)

func TestJoin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		sep      string
		expected string
	}{
		{name: "0 input", input: []string{}, sep: "/", expected: ""},
		{name: "1 input", input: []string{"1"}, sep: "/", expected: "1"},
		{name: "2 inputs", input: []string{"1", "2"}, sep: "/", expected: "1/2"},
		{name: "2 inputs", input: []string{"1", "2"}, sep: ".", expected: "1.2"},
		{name: "3 inputs", input: []string{"1", "2", "3"}, sep: "/", expected: "1/2/3"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := join(testCase.sep, testCase.input...)
			if actual != testCase.expected {
				t.Errorf("input %v and sep %s expects %s but actual %s", testCase.input, testCase.sep, testCase.expected, actual)
			}
		})
	}
}
