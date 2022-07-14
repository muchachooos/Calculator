package main

import (
	"fmt"
	"os"
)

func main() {
	a := "+"
	b := "-"
	c := "/"
	d := "*"
	var j string
	var x float64
	var y float64

	fmt.Print("Что делаем?")
	fmt.Fscan(os.Stdin, &j)

	fmt.Print("Введите первое число.")
	fmt.Fscan(os.Stdin, &x)

	fmt.Print("Введите второе число.")
	fmt.Fscan(os.Stdin, &y)

	if j == a {
		fmt.Println(x + y)
	}

	if j == b {
		fmt.Println(x - y)
	}

	if j == c {
		fmt.Println(x / y)
	}

	if j == d {
		fmt.Println(x * y)
	}
}
