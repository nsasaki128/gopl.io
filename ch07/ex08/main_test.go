package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestLimitReader(t *testing.T) {
	testCases := []struct {
		name     string
		input    multiSort
		expected []*Track
	}{
		{name: "empty", input: multiSort{t: nil}, expected: nil},
		{
			name: "1 sort Title",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byTitle(),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{name: "1 sort Artist",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byArtist(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Album",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byAlbum(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Year",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byYear(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Length",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byLength(),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{name: "2 sort Year and Title",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byYear().byTitle(),
			expected: []*Track{
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{
			name: "2 sort Artist and Title",
			input: multiSort{t: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}}.byArtist().byTitle(),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sort.Sort(testCase.input)
			for i, actual := range testCase.input.t {
				if !reflect.DeepEqual(actual, testCase.expected[i]) {
					t.Errorf("Sort %v results expects %v but actual is %v", testCase.input, testCase.expected, testCase.input.t)
				}
			}
		})
	}
}
