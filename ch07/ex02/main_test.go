package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int64
	}{
		{name: "empty", input: []string{""}, expected: 0},
		{name: "1 word", input: []string{"test"}, expected: 4},
		{name: "2 words", input: []string{"test", "hello", "world"}, expected: 14},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			cw, actual := CountingWriter(&buf)
			for _, s := range testCase.input {
				bs := []byte(s)
				cw.Write(bs)
			}
			if *actual != testCase.expected {
				t.Errorf("<%s> expected %d but actual is %d\n", testCase.input, testCase.expected, actual)
			}
		})
	}
}
