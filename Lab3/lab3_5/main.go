package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	slice = append(slice, 6)
	fmt.Println("Срез после добавления элемента:", slice)
	slice = append(slice[:3], slice[4:]...)
	fmt.Println("Срез после удаления элемента:", slice)
}
