package main

import (
	"testing"
)

func TestLimitReader(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "empty", input: "", expected: true},
		{name: "1 word", input: "a", expected: true},
		{name: "2 word not palindrome", input: "ab", expected: false},
		{name: "2 word palindrome", input: "aa", expected: true},
		{name: "3 word not palindrome", input: "abb", expected: false},
		{name: "3 word palindrome", input: "aba", expected: true},
		{name: "palindrome", input: "nomelonnolemon", expected: true},
		{name: "not palindrome", input: "nomelonnoolemon", expected: false},
		{name: "Japanese palindrome", input: "たけやぶやけた", expected: true},
		{name: "Japanese no palindrome", input: "たけやぶがやけた", expected: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsPalindrome(runeSort([]rune(testCase.input)))
			if actual != testCase.expected {
				t.Errorf("%q expects as palindrome: %v but actual is %v", testCase.input, testCase.expected, actual)
			}
		})
	}
}
