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
	tconn := conn.(*net.TCPConn)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tconn)
		log.Println("done")
		done <- struct{}{} //メインゴルーチンへ通知
	}()
	mustCopy(tconn, os.Stdin)
	tconn.CloseWrite()
	<-done //バックグラウンドのゴルーチンが完了するのを待つ
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
