package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"strings"
)

const (
	resultFile = "fetchResult.txt"
	repeatTime = 1
	urlHeader = "http://"
)


func main () {
	file, err := os.Create(resultFile)
	if err != nil {
		fmt.Println(err)
	}

	writeText := ""
	for i := 0; i < repeatTime; i++ {
		writeText += fetchallTime(os.Args[1:])
	}
	file.WriteString(writeText)

	file.Close()
}

func fetchallTime(urls []string) string {
	outputString := ""
	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(addUrlHeaderIfNeeded(url), ch) // ゴルーチンを開始
	}

	for range urls {
		outputString += (<-ch) // ch チャネルから受信
	}

	totalResult := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	outputString += totalResult

	return outputString
}

func fetch(url string, ch chan<-string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //資源をリークさせない

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

func addUrlHeaderIfNeeded(url string) string {
	if !strings.HasPrefix(url, urlHeader) {
		url = urlHeader + url
	}
	return url
}

