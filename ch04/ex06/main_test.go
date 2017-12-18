package main

import "testing"

func TestCompress(t *testing.T)  {
	testCases := []struct {
			name     string
			input    []byte
			expected []byte
		}{
			{name:"Em space only",input:[]byte("こんにちは　　　世界"), expected:[]byte("こんにちは 世界")},
			{name:"half space only",input:[]byte("こんにちは   世界"), expected:[]byte("こんにちは 世界")},
			{name:"Em and half space",input:[]byte("こんにちは　 　世界"), expected:[]byte("こんにちは 世界")},
		}
		for _, testCase := range testCases {
			actual := compress(testCase.input)
			if len(actual) != len(testCase.expected) {
				t.Errorf("error in case %s \nexpected length:%d\nactual length:\t%d\n", testCase.name, len(testCase.expected), len(actual))
				continue
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != testCase.expected[i] {
					t.Errorf("error in case %s \nexpected:\t%s\nactual:\t%s\n", testCase.name, string(testCase.expected), string(actual))
					break
				}
			}
		}

}
