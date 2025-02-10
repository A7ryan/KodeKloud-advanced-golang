package main

import "fmt"

func calculate(a int, b int) []float64 {
	// your code goes here
	add := float64(a + b)
	sub := float64(a - b)
	mul := float64(a * b)
	div := float64(a / b)
	results := []float64{add, sub, mul, div}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
