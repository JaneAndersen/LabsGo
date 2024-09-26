package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Println("Введите три целых числа")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)
	Average(num1, num2, num3)
}

func Average(num1, num2, num3 int) {
	var avg int = (num1 + num2 + num3) / 3
	fmt.Println("Среднее арифметическое = ", avg)
}
