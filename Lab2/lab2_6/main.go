package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Введите два числа")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	Avg(num1, num2)
}

func Avg(num1, num2 int) {
	var avg int = (num1 + num2) / 2
	fmt.Println("Среднее арифметическое = ", avg)
}
