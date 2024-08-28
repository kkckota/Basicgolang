package main

import (
	"fmt"
	"mygolang/simplecalc"
)

func main() {

	simplecalc.Printmenu()

	var choice int
	fmt.Printf("choose between 1-4\n")
	fmt.Scan(&choice)

	var a, b int
	fmt.Printf("enter the value of a\n")
	fmt.Scan(&a)

	fmt.Printf("enter the value of b\n")
	fmt.Scan(&b)

	switch choice {

	case 1:
		fmt.Printf("Addition %d + %d = %d\n", a, b, simplecalc.Add(a, b))
	case 2:
		fmt.Printf("Subtraction %d - %d = %d\n", a, b, simplecalc.Sub(a, b))
	case 3:
		fmt.Printf("Multiplication %d * %d = %d\n", a, b, simplecalc.Mul(a, b))
	case 4:
		if b == 0 {
			fmt.Printf("cannot be divided by 0")
		} else {
			fmt.Printf("Division %d / %d = %d\n", a, b, simplecalc.Div(a, b))
		}

	}

}
