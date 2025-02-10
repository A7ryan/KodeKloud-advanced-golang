package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	// your code goes here
	switch product {
	case "apples":
		discount := float64(price * 0.10)
		return float64(price - discount)
	case "orange":
		discount := float64(price * 0.20)
		return float64(price - discount)
	case "bananas":
		discount := float64(price * 0) // No discount
		return float64(price - discount)
	default:
		return float64(price)
	}
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
