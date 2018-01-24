package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestLimitReader(t *testing.T) {
	testCases := []struct {
		name        string
		inputString string
		inputN      int64
		expected    string
	}{
		{name: "empty", inputString: "", inputN: 1, expected: ""},
		{name: "word is less than limit", inputString: "hello", inputN: 10, expected: "hello"},
		{name: "word is more than limit", inputString: "hello", inputN: 3, expected: "hel"},
		{name: "boudary testing: word is less than limit", inputString: "hello", inputN: 6, expected: "hello"},
		{name: "boudary testing: word is same as limit", inputString: "hello", inputN: 5, expected: "hello"},
		{name: "boudary testing: word is more than limit", inputString: "hello", inputN: 4, expected: "hell"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(testCase.inputString))
			actual, _ := ioutil.ReadAll(LimitReader(reader, testCase.inputN))
			if string(actual) != testCase.expected {
				t.Errorf("%s for limit %d expects %s but actual is %s", testCase.inputString, testCase.inputN, testCase.expected, string(actual))
			}
		})
	}
}
