package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Введите два целых числа")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	IntOperations(num1, num2)
}

func IntOperations(num1, num2 int) {
	var (
		sum  int = num1 + num2
		ras  int = num1 - num2
		mult int = num1 * num2
	)
	fmt.Println("Сумма чисел = ", sum)
	fmt.Println("Разность чисел = ", ras)
	fmt.Println("Произведение чисел = ", mult)
	if num2 == 0 {
		fmt.Println("Ошибка! Нельзя делить на ноль!")
	} else {
		var del int = num1 / num2
		fmt.Println("Частное чисел = ", del)
	}
}
