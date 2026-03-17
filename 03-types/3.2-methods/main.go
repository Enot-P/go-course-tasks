package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s (%d pages)", b.Title, b.Pages)
}

func main() {
	// task 3.2.1
	fmt.Printf("%s", Book{Title: "title", Pages: 12}.Summary()) // title (12 pages)
}
