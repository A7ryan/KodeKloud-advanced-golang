package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
	// your code goes here
	var result []Item
	for _, item := range items {
		if item.Price >= minPrice && item.Price <= maxPrice {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	items := []Item{
		{Name: "Apple", Price: 0.5},
		{Name: "Banana", Price: 0.25},
		{Name: "Orange", Price: 0.75},
		{Name: "Pineapple", Price: 1.5},
	}

	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
}
