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
			args:[]string{""},
			expected:"",
		},
			{
			name:"one input",
			args:[]string{"hoge"},
			expected:"hoge",
		},
			{
			name:"2 inputs",
			args:[]string{"hoge", "fuga"},
			expected:"hoge fuga",
		},
			{
			name:"several inputs",
			args:[]string{"hoge", "fuga", "piyo", "hogehoge"},
			expected:"hoge fuga piyo hogehoge",
		},
	}

	for _, testCase := range testCases{
		actual := echo(testCase.args)
		if actual != testCase.expected {
			t.Error("case %s exspects %s but actual %s", testCase.name, testCase.expected, actual)
		}
	}
}