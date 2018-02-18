package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	//最初のディレクトリを決める。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//ファイルツリーを走査する。
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDir(root, &n, fileSizes)
		}
		go func() {
			n.Wait()
			close(fileSizes)
		}()
	}()
	//定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	//結果を表示する。
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %1.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
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
