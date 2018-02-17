package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	listen := make(chan string)
	abort := make(chan struct{})
	timer := time.NewTimer(10 * time.Second)

	var wg sync.WaitGroup
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			listen <- input.Text()
		}
		abort <- struct{}{}
	}()
	for {
		select {
		case text := <-listen:
			wg.Add(1)
			go func(s string) {
				echo(c, s, 1*time.Second)
				wg.Done()
			}(text)
			timer.Reset(10 * time.Second)
		case <-timer.C:
			wg.Wait()
			c.Close()
			return
		case <-abort:
			wg.Wait()
			c.Close()
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
