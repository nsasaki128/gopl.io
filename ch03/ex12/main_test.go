package main

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{name: "empty", input1: "", input2: "", expected: true},
		{name: "1 char same", input1: "1", input2: "1", expected: true},
		{name: "1 char different", input1: "1", input2: "2", expected: false},
		{name: "dif length", input1: "11", input2: "1", expected: false},
		{name: "normal anagram sentence", input1: "tom marvolo riddle", input2: "iam lord voldemort", expected: true},
		{name: "normal anagram sentence with Japanese", input1: "Hello, 世界", input2: "世界, Hello", expected: true},
		{name: "non anagram sentence with Japanese", input1: "Hello, 世界", input2: "世界、 Hello", expected: false},
	}
	for _, testCase := range testCases {
		actual := isAnagram(testCase.input1, testCase.input2)
		if actual != testCase.expected {
			t.Errorf("error in case %s \nexpected:\t%t\nactual:\t%t\n", testCase.name, testCase.expected, actual)
		}
	}

}
