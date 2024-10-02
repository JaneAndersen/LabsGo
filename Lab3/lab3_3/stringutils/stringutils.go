package stringutils

import "fmt"

func ReverseString(str string) {
	res := ""
	for i := len(str) - 1; i > -1; i-- {
		res += string(str[i])
	}
	fmt.Println(string(res))
}
