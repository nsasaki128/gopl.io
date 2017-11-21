package main

import (
	"testing"
	"bytes"
)

func TestPrintConvValues(t *testing.T) {
	var testCases = []struct {
		name     string
		in       string
		expected string
	}{
		{name: "0", in: "0",
			expected: "0.00°F = -17.78°C\n0.00°C = 32.00°F\n0.00ft = 0.00m\n0.00m = 0.00ft\n0.00lbs = 0.00kg\n0.00kg = 0.00lbs\n\n"},
		{name: "1", in: "1",
			expected: "1.00°F = -17.22°C\n1.00°C = 33.80°F\n1.00ft = 0.30m\n1.00m = 3.28ft\n1.00lbs = 0.45kg\n1.00kg = 2.20lbs\n\n"},
	}
	for _, testCase := range testCases {
		result := new(bytes.Buffer)
		printConvValues(testCase.in, result)
		if result.String() != testCase.expected {
			t.Errorf("case %s expected %s actual %s", testCase.name, testCase.expected, result)
		}
	}
}
