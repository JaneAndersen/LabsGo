package main

import (
	"fmt"
)

func main() {
	var (
		a, b int
		r    Rectangle
	)
	fmt.Println("Введите длину и ширину прямоугольника")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	r = Rectangle{a, b}
	Square(r)
}

type Rectangle struct {
	a int
	b int
}

func Square(r Rectangle) {
	var s = r.a * r.b
	fmt.Println("Площадь прямоугольника ", s)
}
