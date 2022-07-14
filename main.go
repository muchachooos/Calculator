package main

import (
	"fmt"
	"os"
)

func main() {
	var j string
	var x float64
	var y float64

	fmt.Println("Что делаем?")
	fmt.Fscan(os.Stdin, &j)

	fmt.Println("Введите первое число.")
	fmt.Fscan(os.Stdin, &x)

	fmt.Println("Введите второе число.")
	fmt.Fscan(os.Stdin, &y)

	switch j {
	case "-":
		fmt.Println("Ответ:", x-y)

	case "+":
		fmt.Println("Ответ:", x+y)

	case "*":
		fmt.Println("Ответ:", x*y)

	case "/":
		fmt.Println("Ответ:", x/y)

	default:
		fmt.Println("Но но но.")
	}
}
