package main

import "testing"

func TestTree_String(t *testing.T) {
	testCases := []struct {
		name     string
		input    *tree
		expected string
	}{
		{name: "empty", input: nil, expected: ""},
		{name: "1 value", input: &tree{value: 1, left: nil, right: nil}, expected: "1"},
		{name: "tree", input: &tree{value: 1, left: &tree{value: 0, left: &tree{value: -1, left: nil, right: nil}}, right: &tree{value: 3, left: &tree{value: 2, left: nil, right: nil}, right: nil}}, expected: "-1 0 1 2 3"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.input.String() != testCase.expected {
				t.Errorf("%v expects %s but actual is %s", testCase.input, testCase.expected, testCase.input.String())
			}
		})
	}
}
