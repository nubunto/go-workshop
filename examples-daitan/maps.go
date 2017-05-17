package main

import "fmt"

func main() {
	m := map[string]string{
		"key": "value",
		"b":   "c",
	}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
}
