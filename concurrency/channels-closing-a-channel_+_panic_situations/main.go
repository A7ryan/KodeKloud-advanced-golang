package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)
	val, ok = <-ch
	fmt.Println(val, ok)

	// sending values after closing channel causes panic situation
	// ch <- 1
	// closing already closed channel give panic situation
	// close(ch)
}
