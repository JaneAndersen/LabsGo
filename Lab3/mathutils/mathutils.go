package mathutils

func Factorial(num int) int {
	if num < 1 {
		return 1
	} else {
		return num * Factorial(num-1)
	}
}
