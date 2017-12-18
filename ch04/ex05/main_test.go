package main

import "testing"

func TestDeleteSucDup(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name:"no duplicated", input: []string{"hoge", "fuga", "piyo"}, expected:[]string{"hoge", "fuga", "piyo"}},
		{name:"first duplicated", input: []string{"hoge", "hoge", "fuga", "piyo"}, expected:[]string{"hoge", "fuga", "piyo"}},
		{name:"middle duplicated", input: []string{"hoge", "fuga", "fuga", "piyo"}, expected:[]string{"hoge", "fuga", "piyo"}},
		{name:"last duplicated", input: []string{"hoge", "fuga", "piyo", "piyo"}, expected:[]string{"hoge", "fuga", "piyo"}},
		{name:"duplicated but non-succeeded", input: []string{"hoge", "fuga", "piyo", "fuga", "hoge"}, expected:[]string{"hoge", "fuga", "piyo", "fuga", "hoge"}},
		{name:"duplicated both succeeded and non-succeeded", input: []string{"hoge", "hoge", "fuga", "fuga", "piyo","fuga", "fuga", "hoge"}, expected:[]string{"hoge", "fuga", "piyo", "fuga", "hoge"}},
		{name:"duplicated both succeeded and non-succeeded include Japanese", input: []string{"hoge", "hoge", "fuga", "fuga", "ぴよピヨ","ぴよピヨ", "fuga", "hoge"}, expected:[]string{"hoge", "fuga", "ぴよピヨ", "fuga", "hoge"}},
	}

	for _, testCase := range testCases {
		actual := deleteSucDup(testCase.input)
		if len(actual) != len(testCase.expected) {
			t.Errorf("error in case %s \nexpected length:%d\nactual length:\t%d\n", testCase.name, len(testCase.expected), len(actual))
			continue
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != testCase.expected[i] {
				t.Errorf("error in case %s \nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected, actual)
			}
		}
	}
}