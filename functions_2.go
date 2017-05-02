package main

// START OMIT

import "fmt"

// functions can have multiple return values
func nameAndAge() (string, int) {
	return "John Smith", 36
}

func printMyName(name string) {
	fmt.Println("this is my name:", name)
}

func printMyAge(age int) {
	fmt.Println("this is my age:", age)
}

func main() {
	name, age := nameAndAge()
	printMyName(name)
	printMyAge(age)
}

// END OMIT
