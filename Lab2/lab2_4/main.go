package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("Введите строку (без пробелов)")
	fmt.Scanln(&s)
	Length(s)
}

func Length(s string) {
	var l = len([]rune(s))
	fmt.Println("Количество символов в строке: ", l)
}
