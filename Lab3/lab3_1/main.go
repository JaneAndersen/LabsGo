package main

import (
	"fmt"
	"mathutils"
)

func main() {
	var (
		num int
		f   int
	)
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	f = mathutils.Factorial(num)
	fmt.Println(f)
}
