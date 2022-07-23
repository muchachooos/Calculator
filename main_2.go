package main

import "fmt"

type element struct {
	value       int
	nextElement *element
}

func main() {
	element4 := element{
		4,
		nil,
	}

	element3 := element{
		3,
		&element4,
	}

	element2 := element{
		2,
		&element3,
	}

	element1 := element{
		1,
		&element2,
	}
	El(&element1)
}

func El(*element) {
	fmt.Println()
}
