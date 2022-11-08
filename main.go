package main

import (
	"fmt"
	"os"
)

func main() {
	var operation string
	var first int
	var second int

	fmt.Println("Enter number, please.")
	fmt.Fscan(os.Stdin, &first)

	for {
		fmt.Println("Enter operation, please.")
		fmt.Fscan(os.Stdin, &operation)

		if operation == "=" {
			fmt.Println(first)
			return
		}

		fmt.Println("Enter next number, please.")
		fmt.Fscan(os.Stdin, &second)

		switch operation {
		case "-":
			first -= second
		case "+":
			first += second
		case "*":
			first *= second
		case "/":
			first /= second
		default:
			panic("No no no.")
		}
	}
}
