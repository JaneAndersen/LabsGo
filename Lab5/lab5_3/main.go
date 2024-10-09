package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		r, a, b float64
	)
	fmt.Println("Введите радиус круга")
	fmt.Scanln(&r)
	var c = Circle{r}
	fmt.Println("Площадь круга = ", c.Area())
	fmt.Println("Введите длину прямоугольника")
	fmt.Scanln(&a)
	fmt.Println("Введите ширину прямоугольника")
	fmt.Scanln(&b)
	var rec = Rectangle{a, b}
	fmt.Println("Площадь прямоугольника = ", rec.Area())
	var shapes = []Shape{c, rec}
	Square(shapes)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	a float64
	b float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func Square(s []Shape) {
	for _, v := range s {
		fmt.Println("Площадь фигуры = ", v.Area())
	}
}
