package main

import (
	"fmt"
	"time"
)

// START OMIT
func foo() {
	defer fmt.Println("I run after foo finishes!")
	fmt.Println("I run immediately")
	time.Sleep(1 * time.Second)
}

// END OMIT
func main() {
	foo()
}
