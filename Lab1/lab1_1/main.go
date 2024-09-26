package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("current date + time:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))
}
