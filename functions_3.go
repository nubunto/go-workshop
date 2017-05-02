package main

import "fmt"

// START OMIT

// you can also have naked returns
// discretion is advised, since this hinders readability in larger functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	x, y := split(100)
	fmt.Printf("x: %d -- y: %d", x, y)
}

// END OMIT
