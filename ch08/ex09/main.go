package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

type diskUsage struct {
	nfiles int64
	nbytes int64
}

type dirInfo struct {
	root     string
	fileSize int64
}

func main() {
	//最初のディレクトリを決める。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	usages := map[string]*diskUsage{}
	for _, root := range roots {
		usages[root] = &diskUsage{nfiles: 0, nbytes: 0}
	}

	//ファイルツリーを走査する。
	search := make(chan dirInfo)
	var n sync.WaitGroup
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDir(root, root, &n, search)
		}
		go func() {
			n.Wait()
			close(search)
		}()
	}()
	//定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	//結果を表示する。
loop:
	for {
		select {
		case s, ok := <-search:
			if !ok {
				break loop
			}
			usages[s.root].nfiles++
			usages[s.root].nbytes += s.fileSize
		case <-tick:
			printDiskUsage(usages)
		}
	}
	printDiskUsage(usages)
}

func printDiskUsage(usages map[string]*diskUsage) {
	sorted := []string{}
	for k := range usages {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for _, key := range sorted {
		fmt.Printf("%s %d files %1.1f GB\n", key, usages[key].nfiles, float64(usages[key].nbytes)/1e9)
	}
}

func walkDir(root, dir string, n *sync.WaitGroup, search chan<- dirInfo) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, search)
		} else {
			search <- dirInfo{root: root, fileSize: entry.Size()}
		}
	}
}

//semaは、direntsでの並行性を制限するための係数セマフォです。
var sema = make(chan struct{}, 20)

//dirents haはディレクトリdirの項目を返します。
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //tokenを獲得
	defer func() { <-sema }() //tokenを解放
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
