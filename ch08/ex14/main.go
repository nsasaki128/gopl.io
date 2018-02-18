package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string //送信用メッセージチャネル
type enter struct {
	cli  client
	name string
}

var (
	entering = make(chan enter)
	leaving  = make(chan enter)
	messages = make(chan string) //クライアントから受信する全てのメッセージ
)

func broadcaster() {
	clients := make(map[enter]bool) //全ての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			// 受信するメッセージを全てのクライアントの
			// 送信用メッセージチャネルヘブロードキャストする。
			for cli := range clients {
				cli.cli <- msg
			}
		case cli := <-entering:
			names := []string{}
			for cli := range clients {
				names = append(names, cli.name)
			}
			sort.Strings(names)
			cli.cli <- "current clients: " + strings.Join(names, ", ")

			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) //送信用のクライアントメッセージ
	go clientWriter(conn, ch)
	ch <- "Who are you?"
	who := string("")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		who = input.Text()
		break
	}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- enter{cli: ch, name: who}

	inputs := make(chan string)
	go func() {
		for input.Scan() {
			inputs <- input.Text()
		}
	}()

loop:
	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Fprintln(conn, "Bye")
			conn.Close()
			break loop
		case text := <-inputs:
			messages <- who + ": " + text
		}
	}
	//注意: input.Err()からの潜在的なエラーを無視している

	leaving <- enter{cli: ch, name: who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
