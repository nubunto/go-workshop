package main

import "fmt"

func main() {
	sliceOfStrings := []string{"a", "b", "c"}
	fmt.Println(sliceOfStrings)

	for i, s := range sliceOfStrings {
		fmt.Println("#", i, ":", s)
	}

	fmt.Println(len(sliceOfStrings))
	sliceOfStrings = append(sliceOfStrings, "d")

	for i, s := range sliceOfStrings {
		fmt.Println("#", i, ":", s)
	}

	fmt.Println(len(sliceOfStrings))
}
