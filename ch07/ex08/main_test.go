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
			name: "no sort Title",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{
			name: "1 sort Title",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byTitle(),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{name: "1 sort Artist",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byArtist(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Album",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byAlbum(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Year",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byYear(),
			expected: []*Track{
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		{name: "1 sort Length",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byLength(),
			expected: []*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{name: "2 sort Year and Title",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byYear().byTitle(),
			expected: []*Track{
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			},
		},
		{
			name: "2 sort Artist and Title",
			input: newMultiSort([]*Track{
				{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
				{"GO", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
			}).byArtist().byTitle(),
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

func newData() []*Track {
	return []*Track{
		{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
		{"GO", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

type testSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (m testSort) Len() int           { return len(m.t) }
func (m testSort) Less(i, j int) bool { return m.less(m.t[i], m.t[j]) }
func (m testSort) Swap(i, j int)      { m.t[i], m.t[j] = m.t[j], m.t[i] }

func BenchmarkStableSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := newData()
		sort.Stable(testSort{t: s, less: func(x, y *Track) bool { return x.Title < y.Title }})
		sort.Stable(testSort{t: s, less: func(x, y *Track) bool { return x.Artist < y.Artist }})
		sort.Stable(testSort{t: s, less: func(x, y *Track) bool { return x.Album < y.Album }})
		sort.Stable(testSort{t: s, less: func(x, y *Track) bool { return x.Year < y.Year }})
		sort.Stable(testSort{t: s, less: func(x, y *Track) bool { return x.Length < y.Length }})
	}
}

func BenchmarkMultiSort(b *testing.B) {
	s := newMultiSort(newData()).byTitle().byArtist().byAlbum().byYear().byLength()
	sort.Stable(s)
}
