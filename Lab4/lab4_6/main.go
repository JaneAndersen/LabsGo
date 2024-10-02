package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)
	fmt.Println("Введите длину массива")
	fmt.Scanln(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Введите число")
		fmt.Scanln(&num[i])
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Print(num[i])
	}
}
