package main

import (
	"fmt"
	"stringutils"
)

func main() {
	var (
		s string
	)
	fmt.Println("Введите строку")
	fmt.Scanln(&s)
	stringutils.ReverseString(s)
}
