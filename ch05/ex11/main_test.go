package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string][]string
		isCyclic bool
	}{
		{name: "1 straight input", input: map[string][]string{"1": {"2"}}, isCyclic: false},
		{name: "1 cyclic input", input: map[string][]string{"1": {"1"}}, isCyclic: true},
		{name: "2 cyclic inputs", input: map[string][]string{"1": {"2"}, "2": {"1"}}, isCyclic: true},
		{name: "long cyclic inputs", input: map[string][]string{"1": {"2", "3", "4"}, "2": {"3", "4"}, "3": {"4"}, "4": {"5", "6"}, "5": {"1"}}, isCyclic: true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := topoSort(testCase.input)
			if (err != nil) != testCase.isCyclic {
				t.Errorf("cyclic expects %v but actual %v", testCase.isCyclic, err != nil)
			}
		})
	}
}
