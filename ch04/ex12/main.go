package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	BaseURLDomain  = "https://xkcd.com/"
	BaseURLJson    = "info.0.json"
	MaxNotFoundNum = 5
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

var maxNum = flag.Int("n", 100, "number of comics")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s, search words\n", os.Args[0])
		os.Exit(1)
	}

	var comics []*Comic
	comics = getComics()
	for _, word := range args {
		search(word, comics)
	}
}

func search(word string, comics []*Comic) {
	fmt.Println("===================================")
	fmt.Println("Start searching comic information...")
	fmt.Printf("The word is %s\n", word)
	for i, comic := range comics {
		if comic == nil {
			continue
		}
		if isComicContainsWord(word, *comic) {
			fmt.Printf("\nFind the word %s!!\n", word)
			fmt.Printf("URL\n%s\n", comicUrl(i+1))
			fmt.Printf("TRANSCIRPT\n%s\n", comic.Transcript)
		}
	}
	fmt.Println("\nFinish searching comic information!!!")
	fmt.Println("===================================")

}

func isComicContainsWord(word string, comic Comic) bool {
	return strings.Contains(comic.Transcript, word)
}

func getComics() []*Comic {
	i := 1
	notFoundNum := 0
	var comics []*Comic
	fmt.Println("Start gathering comic information...")
	for {
		if i > *maxNum || notFoundNum > MaxNotFoundNum {
			break
		}
		fmt.Print(".")
		c, err := getComic(comicUrl(i))
		if err != nil {
			notFoundNum++
		}
		comics = append(comics, c)
		i++
	}
	fmt.Println("\nFinish gathering comic information!!!")
	return comics
}

func getComic(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var comic *Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return comic, nil
}

func comicUrl(i int) string {
	return BaseURLDomain + fmt.Sprint(i) + "/" + BaseURLJson
}
