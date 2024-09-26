package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	fmt.Println("Введите два числа с плавающей запятой")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println("")
	FlOperations(num1, num2)
}

func FlOperations(num1, num2 float64) {
	var (
		sum float64 = num1 + num2
		ras float64 = num1 - num2
	)
	fmt.Println("Сумма чисел = ", sum)
	fmt.Println("Разность чисел = ", ras)
}
