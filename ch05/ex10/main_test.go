package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	testCases := []struct {
		name  string
		input map[string]map[string]bool
	}{
		{name: "prereqs", input: prereqs},
		{name: "1 input", input: map[string]map[string]bool{"2": {"1": true}}},
		{name: "2 inputs", input: map[string]map[string]bool{"2": {"1": true}, "3": {"2": true}}},
		{name: "diamond shape inputs", input: map[string]map[string]bool{"2": {"1": true}, "3": {"1": true}, "4": {"2": true, "3": true}}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			found := make(map[string]bool)
			for _, course := range topoSort(testCase.input) {
				for before := range testCase.input[course] {
					if !found[before] {
						t.Errorf("%s should appear before %s", before, course)
					}
				}
				found[course] = true
			}
		})
	}
}
