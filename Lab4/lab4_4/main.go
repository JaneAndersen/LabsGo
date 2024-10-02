package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s string
	)
	fmt.Println("")
	fmt.Scanln(&s)
	fmt.Println(strings.ToUpper(s))
}
