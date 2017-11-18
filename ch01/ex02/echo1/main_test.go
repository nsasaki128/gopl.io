package main

import "testing"

func TestEcho(t *testing.T) {
	var testCases = []struct {
		name string
		args []string
		expected string
	}{
		{
			name:"empty",
			args:[]string{},
			expected:"",
		},
			{
			name:"one input",
			args:[]string{"hoge"},
			expected:"1 hoge\n",
		},
			{
			name:"2 inputs",
			args:[]string{"hoge", "fuga"},
			expected:"1 hoge\n2 fuga\n",
		},
			{
			name:"several inputs",
			args:[]string{"hoge", "fuga", "piyo", "hogehoge"},
			expected:"1 hoge\n2 fuga\n3 piyo\n4 hogehoge\n",
		},
	}

	for _, testCase := range testCases{
		actual := echoWithIndex(testCase.args)
		if actual != testCase.expected {
			t.Error("case %s exspects %s but actual %s", testCase.name, testCase.expected, actual)
		}
	}
}