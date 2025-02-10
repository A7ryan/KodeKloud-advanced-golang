package main

import "fmt"

var activeUserCount = 0

// Declare variable activeUserCount
// your code goes here

func entry() {
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here
	activeUserCount++
}

func exit() {
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
