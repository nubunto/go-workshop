package main

import "fmt"

// START OMIT

// functions are declared with the "func" keyword
// note how the identifier name and type is reversed from C/Java
// void printMyName(String name) {
func printMyName(name string) {
	fmt.Println("say my name:", name)
}

func main() {
	printMyName("Heisenberg")
}

// END OMIT
