package main

import (
	"fmt"
	"os"
	"sort"
)

// prereqsは情報科学の各講座をそれぞれの事前条件となる講座と対応付けします。
var prereqs = map[string][]string{
	"algorithms":      {"data structures"},
	"calculus":        {"liniear algebra"},
	"liniear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input %v\n", err)
		os.Exit(1)
	}
	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	temp := make(map[string]bool)
	isCyclic := false
	var cyclicItem string
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if isCyclic {
				return
			}
			if temp[item] {
				isCyclic = true
				cyclicItem = item
				return
			}
			if !seen[item] {
				seen[item] = true

				temp[item] = true
				visitAll(m[item])
				temp[item] = false

				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	if isCyclic {
		return nil, fmt.Errorf("cyclic topology find for item %s", cyclicItem)
	}
	return order, nil
}
