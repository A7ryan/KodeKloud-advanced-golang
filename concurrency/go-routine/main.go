package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(1 * time.Second)
}

func start() {
	go process()
	fmt.Println("Started..")
}

func process() {
	fmt.Println("In Process..")
}
