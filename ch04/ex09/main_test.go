package main

import (
	"testing"
	"strings"
	"reflect"
)

func TestWordfreq(t *testing.T){
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{name:"only English different words", input:"hoge fuga piyo", expected:map[string]int{"hoge":1, "fuga":1, "piyo":1}},
		{name:"only Japanese different words", input:"ほげ ふが ピヨ", expected:map[string]int{"ほげ":1, "ふが":1, "ピヨ":1}},
		{name:"English and Japanese many words", input:"hoge 世界 hoge fuga ほげ fugafuga ふが ピヨ", expected:map[string]int{"hoge":2, "fuga":1, "fugafuga":1, "世界":1, "ほげ":1, "ふが":1, "ピヨ":1}},
	}
	for _, testCase := range testCases {
		actual := wordfreq(strings.NewReader(testCase.input))
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("error in case %s \nexpected:\t%v\nactual:\t%v\n", testCase.name, testCase.expected, actual)
		}
	}


}
