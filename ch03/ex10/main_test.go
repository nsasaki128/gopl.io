package main

import "testing"

func TestComma(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected:""},
		{name: "len 1", input: "1", expected:"1"},
		{name: "len 2", input: "12", expected:"12"},
		{name: "len 3", input: "123", expected:"123"},
		{name: "len 4", input: "1234", expected:"1,234"},
		{name: "len 5", input: "12345", expected:"12,345"},
		{name: "len 6", input: "123456", expected:"123,456"},
		{name: "len 7", input: "1234567", expected:"1,234,567"},
	}
	for _, testCase := range testCases {
		actual := comma(testCase.input)
		if actual != testCase.expected {
			t.Errorf("error in case %s \nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected, actual)

		}
	}

}
