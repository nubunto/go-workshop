package main

import "fmt"

func main() {
	// START OMIT
	var x int
	// variables are always initialized by the runtime with their 0 value
	fmt.Println(x) // prints 0

	var y complex128 // yay, complex numbers
	y = 18 + 4i
	fmt.Println(y)
	fmt.Println(y + 1)

	// Go has (very basic) type inference
	z := "fizzbuzz"
	fmt.Printf("this is the string: %s", z)

	// you can declare multiple variables in one go
	i, j, k := "foo", 0, -2.6
	fmt.Println(i, j, k) // "foo", 0, -2.6

	// END OMIT
}
