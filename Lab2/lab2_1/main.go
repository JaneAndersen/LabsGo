package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Println("Введите число")
	fmt.Scanln(&num1)
	if num1%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
