package main

import (
	"fmt"
)

/*
	A slice is a section of an underlying array. Changing elements of a
	slice modifies the corresponding elements of its underlying array.
*/

func main() {
	arr := [4]string{
		"Hello",
		"World",
		"Whats",
		"Up",
	}
	fmt.Println(arr)

	a := arr[0:2]
	b := arr[1:3]
	fmt.Println(a, b)

	b[0] = "wow"
	fmt.Println(a, b)
	fmt.Println(arr)
}
