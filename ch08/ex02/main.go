package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("port", 8000, "listen port (default 8000)")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		ftp := newFtp(conn)
		ftp.run()
		conn.Close()
	}
}
