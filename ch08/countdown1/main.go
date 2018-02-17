package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing condition.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
