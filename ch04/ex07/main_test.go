package main

import "testing"

func TestReverseByte(t *testing.T)  {
	testCases := []struct {
			name     string
			input    string
			expected string
		}{
			{name:"Japanese only",input:"こんにちは、　世界", expected:"界世　、はちにんこ"},
			{name:"English only",input:"Hello, world", expected:"dlrow ,olleH"},
			{name:"Both Japanese and English",input:"Hello, 世界", expected:"界世 ,olleH"},
		}
		for _, testCase := range testCases {
			byteInput := []byte(testCase.input)
			reverseByte(byteInput)
			actual := string(byteInput)
			if actual != testCase.expected {
				t.Errorf("error in case %s \nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected, actual)
			}
		}

}
