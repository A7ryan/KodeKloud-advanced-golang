package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	words := strings.Fields(text)
	frequecy := make(map[string]int)
	for _, word := range words {
		frequecy[word]++
	}
	return frequecy
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}
