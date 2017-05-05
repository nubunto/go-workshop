package main

import "fmt"

func main() {
	// START OMIT
	var x int
	// variables are always initialized by the runtime with their zero value (to avoid random garbage, like in C)
	// 0 for numbers, "" for strings, nil for pointers/channels/slices/maps
	// for structs, all the fields have their respective zero values, as above
	fmt.Println(x) // prints 0

	var y complex128 // yay, complex numbers
	y = 18 + 4i
	fmt.Println(y)
	fmt.Println(y + 2i)

	var p float64 = 3.90341 // floats have arbitrary precision
	fmt.Println("my float: %.2f", p)

	// Go has (very basic) type inference
	z := "fizzbuzz"
	fmt.Printf("this is the string: %s\n", z)

	// you can declare multiple variables in one Go (pun intended)
	i, j, k := "foo", 0, -2.6
	fmt.Println(i, j, k) // "foo" 0 -2.6

	// END OMIT
}
