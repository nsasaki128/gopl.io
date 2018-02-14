package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type clockwall struct {
	place string
	time  string
}

func main() {
	c := make(chan *clockwall)
	var places []string
	times := map[string]string{}

	for _, arg := range os.Args[1:] {
		sp := strings.Split(arg, "=")
		if len(sp) != 2 {
			log.Fatalf("invalid parameter: %s", arg)
		}
		place, addr := sp[0], sp[1]
		times[place] = ""
		places = append(places, place)

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		go func(place string, conn net.Conn, c chan<- *clockwall) {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				c <- &clockwall{place: place, time: scanner.Text()}
			}
		}(place, conn, c)
	}

	for pc := range c {
		times[pc.place] = pc.time

		// information about showing clear terminal screen from
		// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		for _, place := range places {
			fmt.Printf("%s\t%s\n", place, times[place])
		}
	}

}
