package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
	book.Pages = pages
}

func main() {

	book1 := &Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Pages: 180}
	book2 := &Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", Pages: 281}
	book3 := &Book{Title: "Pride and Prejudice", Author: "Jane Austen", Pages: 279}

	updatePages(book1, 210)
	fmt.Println(book1)
	updatePages(book2, 250)
	fmt.Println(book2)
	updatePages(book3, 295)
	fmt.Println(book3)

	// your code for printing objects goes here
}
