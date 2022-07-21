package main

import (
	"fmt"
	"os"
)

func main() {
	var operation string
	var first int
	var second int

	fmt.Println("Введите число.")
	fmt.Fscan(os.Stdin, &first)

	for {
		fmt.Println("Введите операцию.")
		fmt.Fscan(os.Stdin, &operation)

		if operation == "=" {
			fmt.Println(first)
			return
		}

		fmt.Println("Введите следующее число.")
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
			panic("Но но но.")
		}
	}
}
