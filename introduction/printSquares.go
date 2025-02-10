package main

import "fmt"

func printSquares(n int) {
	// base case
	if n == -6 {
		return
	}
	fmt.Printf("%d ", n*n)
	printSquares(n - 1)
}

func main() {
	printSquares(2)
}
