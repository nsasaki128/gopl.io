package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		f        func(s string) string
		expected string
	}{
		{name: "$ notfound", input: "testtest", f: func(s string) string { return s + ":" + s }, expected: "testtest"},
		{name: "1 $ repeat function", input: "test$test", f: func(s string) string { return s + ":" + s }, expected: "testtest:test"},
		{name: "1 $ add newline function", input: "test$test", f: func(s string) string { return s + "\n" }, expected: "testtest\n"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := expand(testCase.input, testCase.f)
			if actual != testCase.expected {
				t.Errorf("%s for func %T expects %s but actual %s \n", testCase.input, testCase.f, testCase.expected, actual)
			}
		})
	}
}
