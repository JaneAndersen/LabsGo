package main

import (
	"fmt"
)

func main() {
	var (
		n   int
		num int
		sum int
	)
	sum = 0
	fmt.Println("Сколько чисел вы хотите сложить?")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Введите число")
		fmt.Scanln(&num)
		sum += num
	}
	fmt.Println("Сумма чисел: ", sum)
}
