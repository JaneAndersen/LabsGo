package main

import (
	"fmt"
)

func main() {
	var (
		arr [5]int
	)
	fmt.Println("")
	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}
