package main

import (
	"fmt"

	"LabsGo/Lab3/mathutils"
)

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	mathutils.Factorial(num)
}
