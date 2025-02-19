package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sell(ch)
	go buy(ch)
	time.Sleep(1 * time.Second)
}

// send data to channel
func sell(ch chan string) {
	ch <- "Furniture"
	// it blocks code below this until another go routine receives this value
	fmt.Println("Sending Data")
}

// receive data from channel
func buy(ch chan string) {
	fmt.Println("Waiting for data..")
	val := <-ch
	fmt.Println("Received Data: ", val)
}
