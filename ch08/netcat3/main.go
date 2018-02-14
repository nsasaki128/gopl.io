package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //注意:エラーを無視している
		log.Println("done")
		done <- struct{}{} //メインゴルーチンへ通知
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //バックグラウンドのゴルーチンが完了するのを待つ
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
