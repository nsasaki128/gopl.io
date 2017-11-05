package main

import "testing"

func TestAddUrlHeaderIfNeeded(t *testing.T) {
	var tests = []struct{
		caseName string
		input string
		expectedResult string
	}{
		{
			caseName       : string("Less Header"),
			input          : string("gopl.io"),
			expectedResult : string("http://gopl.io"),
		},
		{
			caseName       : string("Has Header"),
			input          : string("http://gopl.io"),
			expectedResult : string("http://gopl.io"),
		},
		{
			caseName       : string("Empty"),
			input          : string(""),
			expectedResult : string("http://"),
		},
		{
			caseName       : string("Has https Header"),
			input          : string("https://gopl.io"),
			expectedResult : string("http://https://gopl.io"),
		},
		{
			caseName       : string("only ip"),
			input          : string("8.8.8.8"),
			expectedResult : string("http://8.8.8.8"),
		},

	}



	for _, testCase := range tests {
		if result := addUrlHeaderIfNeeded(testCase.input); result != testCase.expectedResult {
			t.Errorf("error in case %s input text %s expects %s but result is %s\n",
				testCase.caseName, testCase.input, testCase.expectedResult, result)
		}
	}

}
