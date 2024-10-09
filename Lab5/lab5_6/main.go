package main

import (
	"fmt"
)

func main() {
	var b Book = Book{"X-wing", "Michael Stackpole", 411}
	b.Info()
}

type Stringer interface {
	Info()
}

type Book struct {
	title  string
	author string
	pages  int
}

func (b Book) Info() {
	fmt.Println("Название книги:\n", b.title, "\nАвтор книги:\n", b.author, "\nКоличество страниц:\n", b.pages)
}
