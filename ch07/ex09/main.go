package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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

var trackList = template.Must(template.New("trackTable").Parse(`
<h1>Go music</h1>
<table>
<tr style='text-align: left'>
  <th><a href='{{.NewURL "Title"}}'>Title</a></th>
  <th><a href='{{.NewURL "Artist"}}'>Artist</a></th>
  <th><a href='{{.NewURL "Album"}}'>Album</a></th>
  <th><a href='{{.NewURL "Year"}}'>Year</a></th>
  <th><a href='{{.NewURL "Length"}}'>Length</a></th>
</tr>
{{range .Tracks}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

type data struct {
	Tracks []*Track
	r      *http.Request
}

func (d *data) NewURL(sortKey string) *url.URL {
	u := *d.r.URL
	q := u.Query()
	q.Add("sort", sortKey)
	u.RawQuery = q.Encode()
	return &u
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sortByQuery(tracks, r.URL.Query())
	trackTable := data{Tracks: tracks, r: r}
	if err := trackList.Execute(w, &trackTable); err != nil {
		log.Fatal(err)
	}
}
func newMultiSort(t []*Track) multiSort {
	s := multiSort{t: t, less: func(x, y *Track) bool { return false }}
	return s
}
func sortByQuery(t []*Track, q url.Values) {
	s := newMultiSort(t)
	for _, key := range q["sort"] {
		switch key {
		case "Title":
			s = s.byTitle()
		case "Artist":
			s = s.byArtist()
		case "Album":
			s = s.byAlbum()
		case "Year":
			s = s.byYear()
		case "Length":
			s = s.byLength()
		}
	}
	sort.Sort(s)
}
