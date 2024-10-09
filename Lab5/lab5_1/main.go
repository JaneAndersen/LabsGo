package main

import (
	"fmt"
)

func main() {
	var p Person
	p = Person{"Вася", 12}
	output(p)
	p.Birthday()
	fmt.Println(p.age)
}

type Person struct {
	name string
	age  int
}

func output(p Person) {
	fmt.Println("Имя: ", p.name, "\nВозраст: ", p.age)
}

func (p *Person) Birthday() {
	p.age++
}
