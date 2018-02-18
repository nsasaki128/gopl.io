package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //クライアントから受信する全てのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) //全ての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			// 受信するメッセージを全てのクライアントの
			// 送信用メッセージチャネルヘブロードキャストする。
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) //送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	//注意: input.Err()からの潜在的なエラーを無視している

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
