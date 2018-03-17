package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		name    string
		input   string
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{name: "empty", input: "", counts: map[rune]int{}, utflen: []int{0, 0, 0, 0, 0}, invalid: 0},
		{name: "one English", input: "a", counts: map[rune]int{'a': 1}, utflen: []int{0, 1, 0, 0, 0}, invalid: 0},
		{name: "one Japanese", input: "あ", counts: map[rune]int{'あ': 1}, utflen: []int{0, 0, 0, 1, 0}, invalid: 0},
		{name: "only English", input: "Hello", counts: map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1}, utflen: []int{0, 5, 0, 0, 0}, invalid: 0},
		{name: "only Japanese", input: "アイウエオ", counts: map[rune]int{'ア': 1, 'イ': 1, 'ウ': 1, 'エ': 1, 'オ': 1}, utflen: []int{0, 0, 0, 5, 0}, invalid: 0},
		{name: "both English and Japanese", input: "Hello 世界", counts: map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1, ' ': 1, '世': 1, '界': 1}, utflen: []int{0, 6, 0, 2, 0}, invalid: 0},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			in := bufio.NewReader(strings.NewReader(testCase.input))
			counts, utflen, invalid, err := charcount(in)
			if !reflect.DeepEqual(counts, testCase.counts) || invalid != testCase.invalid || err != nil {
				t.Errorf("charcount(%q)\nexpect\tcounts:%v, utflen:%v, invalid:%v, err:nil\nactual\tcounts:%v, utflen:%v, invalid:%v, err:%v", testCase.input,
					testCase.counts, testCase.utflen, testCase.invalid,
					counts, utflen, invalid, err)
			}
			for i, len := range utflen {
				if len != testCase.utflen[i] {
					t.Errorf("charcount(%q)\nexpect\tcounts:%v, utflen:%v, invalid:%v, err:nil\nactual\tcounts:%v, utflen:%v, invalid:%v, err:%v", testCase.input,
						testCase.counts, testCase.utflen, testCase.invalid,
						counts, utflen, invalid, err)
				}
			}
		})
	}
}
