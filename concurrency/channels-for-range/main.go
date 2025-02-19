package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println("Sent all data to the channel")
	close(ch)
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data...")
	for val := range ch {
		fmt.Println("Received data: ", val)
	}
	wg.Done()
}
