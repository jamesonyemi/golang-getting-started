package main

import "fmt"

func main() {
	book := "The color of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Println(book[0])
	fmt.Printf("book[0]=%v of (Type %T\n)", book[0], book[0])

	poem := `
		The desert lions of Morocco are
		coming home with the cup
		---
	`
	fmt.Println(poem)

}