package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Println("Введите число")
	fmt.Scanln(&num1)
	PNZ(num1)
}

func PNZ(num int) {
	if num == 0 {
		fmt.Println("Zero")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Positive")
	}
}
