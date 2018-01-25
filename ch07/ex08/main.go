package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"GO", "Dlilah", "From the Roots Up", 2012, length("3m38s")},
	{"GO", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTrakcs(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type multiSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func newMultiSort(t []*Track) multiSort {
	s := multiSort{t: t, less: func(x, y *Track) bool { return false }}
	return s
}

func (m multiSort) Len() int           { return len(m.t) }
func (m multiSort) Less(i, j int) bool { return m.less(m.t[i], m.t[j]) }
func (m multiSort) Swap(i, j int)      { m.t[i], m.t[j] = m.t[j], m.t[i] }

func (m multiSort) byTitle() multiSort {
	return multiSort{m.t, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return m.less(x, y)
	}}
}

func (m multiSort) byArtist() multiSort {
	return multiSort{m.t, func(x, y *Track) bool {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		return m.less(x, y)
	}}
}

func (m multiSort) byAlbum() multiSort {
	return multiSort{m.t, func(x, y *Track) bool {
		if x.Album != y.Album {
			return x.Album < y.Album
		}
		return m.less(x, y)
	}}
}

func (m multiSort) byYear() multiSort {
	return multiSort{m.t, func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return m.less(x, y)
	}}
}

func (m multiSort) byLength() multiSort {
	return multiSort{m.t, func(x, y *Track) bool {
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return m.less(x, y)
	}}
}

func main() {
	printTrakcs(tracks)
	fmt.Println("After sort title year")
	sort.Sort(newMultiSort(tracks).byYear().byTitle())
	printTrakcs(tracks)
	fmt.Println("After sort title artist")
	sort.Sort(newMultiSort(tracks).byArtist().byTitle())
	printTrakcs(tracks)
}
