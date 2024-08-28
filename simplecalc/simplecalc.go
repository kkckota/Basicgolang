package simplecalc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		fmt.Println("Division error")
		return 0
	}
	return a / b
}
