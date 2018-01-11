package main

import (
	"fmt"
)

// prereqsは情報科学の各講座をそれぞれの事前条件となる講座と対応付けします。
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"liniear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for k := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	for k := range m {
		visitAll(map[string]bool{k: true})
	}
	return order
}
