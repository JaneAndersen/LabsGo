package main

import (
	"fmt"
)

func main() {
	s_arr := [5]string{"golang", "java", "c", "assembler", "python"}
	max := ""
	for _, s := range s_arr {
		if len(s) > len(max) {
			max = s
		}
	}
	fmt.Println("Самая длинная строка: ", max)

}
