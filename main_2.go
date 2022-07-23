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

	element1.nextElement = &element2
	element2.nextElement = &element3
	element3.nextElement = &element4

	var c *element = &element1

	for c != nil {
		fmt.Println(c.value)
		c = c.nextElement
	}
}
