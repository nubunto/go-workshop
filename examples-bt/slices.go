package main

import "fmt"

func main() {
	sliceOfStrings := []string{"a", "b", "c"}
	for _, s := range sliceOfStrings {
		fmt.Println(s)
	}

	sliceOfStrings = append(sliceOfStrings, "d")
	for _, s := range sliceOfStrings {
		fmt.Println(s)
	}
}
