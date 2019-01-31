package main

import (
	"fmt"
	"time"
)

func sleeper(label string, ping chan<- string) func() {
	return func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ping <- label
		}
	}
}

func main() {
	ping := make(chan string)
	go sleeper("1", ping)()
	go sleeper("2", ping)()
	go sleeper("3", ping)()
	go sleeper("4", ping)()
	go sleeper("5", ping)()

	for {
		select {
		case p := <-ping:
			fmt.Println(p, "awake!")
		}
	}
}

