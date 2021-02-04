package main

import (
	"fmt"
)

/* Go arrays cannot be dynamically allocated or resized */

func main() {
	var a [10]string

	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11}

	fmt.Println(a)
	fmt.Println(primes)
}
