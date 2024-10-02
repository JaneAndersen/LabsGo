package main

import (
	"fmt"
)

func main() {
	var (
		avg  int
		name string
	)
	var people = map[string]int{
		"Степан":    22,
		"Дмитрий":   25,
		"Дарья":     20,
		"Анастасия": 19,
	}
	people["Александра"] = 23
	for key, value := range people {
		fmt.Println(key, value)
	}
	avg = AverAge(people)
	fmt.Println("", avg)
	fmt.Println("")
	fmt.Scanln(&name)
	delete(people, name)

}

func AverAge(m map[string]int) int {
	var (
		age     int
		counter int
		avg     int
	)
	age = 0
	counter = 0
	for _, v := range m {
		age += v
		counter++
	}
	avg = age / counter
	return avg
}
