package main

import "testing"

func TestComma(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "len 1", input: "1", expected: "1"},
		{name: "len 2", input: "12", expected: "12"},
		{name: "len 3", input: "123", expected: "123"},
		{name: "len 4", input: "1234", expected: "1,234"},
		{name: "len 5", input: "12345", expected: "12,345"},
		{name: "len 6", input: "123456", expected: "123,456"},
		{name: "len 7", input: "1234567", expected: "1,234,567"},
		{name: "fractional len 1", input: "0.1", expected: "0.1"},
		{name: "fractional len 2", input: "0.12", expected: "0.12"},
		{name: "fractional len 3", input: "0.123", expected: "0.123"},
		{name: "fractional len 4", input: "0.1234", expected: "0.1234"},
		{name: "fractional len 5", input: "0.12345", expected: "0.12345"},
		{name: "fractional len 6", input: "0.123456", expected: "0.123456"},
		{name: "fractional len 7", input: "0.1234567", expected: "0.1234567"},
		{name: "minus len 1", input: "-1", expected: "-1"},
		{name: "minus len 2", input: "-12", expected: "-12"},
		{name: "minus len 3", input: "-123", expected: "-123"},
		{name: "minus len 4", input: "-1234", expected: "-1,234"},
		{name: "minus len 5", input: "-12345", expected: "-12,345"},
		{name: "minus len 6", input: "-123456", expected: "-123,456"},
		{name: "minus len 7", input: "-1234567", expected: "-1,234,567"},
		{name: "plus len 1", input: "+1", expected: "+1"},
		{name: "plus len 2", input: "+12", expected: "+12"},
		{name: "plus len 3", input: "+123", expected: "+123"},
		{name: "plus len 4", input: "+1234", expected: "+1,234"},
		{name: "plus len 5", input: "+12345", expected: "+12,345"},
		{name: "plus len 6", input: "+123456", expected: "+123,456"},
		{name: "plus len 7", input: "+1234567", expected: "+1,234,567"},
		{name: "minus with fractional len 1", input: "-1.1", expected: "-1.1"},
		{name: "plus with fractional len 2", input: "+12.12", expected: "+12.12"},
		{name: "int with fractional len 3", input: "123.123", expected: "123.123"},
		{name: "minus with fractional len 4", input: "-1234.1234", expected: "-1,234.1234"},
		{name: "plus with fractional len 5", input: "+12345.12345", expected: "+12,345.12345"},
		{name: "int with fractional len 6", input: "123456.123456", expected: "123,456.123456"},
		{name: "minus with fractional len 7", input: "-1234567.1234567", expected: "-1,234,567.1234567"},
		{name: "plus with fractional len 8", input: "+12345678.12345678", expected: "+12,345,678.12345678"},
		{name: "int with fractional len 9", input: "123456789.123456789", expected: "123,456,789.123456789"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := comma(testCase.input)
			if actual != testCase.expected {
				t.Errorf("error in case %s \nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected, actual)

			}
		})
	}

}
