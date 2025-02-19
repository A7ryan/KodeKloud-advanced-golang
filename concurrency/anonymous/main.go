// don't have any name
// can also be called via go-routine

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hey, I am an anonymous function!")
	}()

	time.Sleep(1 * time.Second)
}
