package main

import "fmt"

type Element struct {
	value       int
	nextElement *Element
}

func main() {
	element4 := Element{
		4,
		nil,
	}

	element3 := Element{
		3,
		&element4,
	}

	element2 := Element{
		2,
		&element3,
	}

	element1 := Element{
		1,
		&element2,
	}
	print(&element1)
}

func print(root *Element) {
	fmt.Println(root.value)
	fmt.Println(root.nextElement.value)
	fmt.Println(root.nextElement.nextElement.value)
	fmt.Println(root.nextElement.nextElement.nextElement.value)
}
