package main

import "fmt"

func main() {
	m := map[string]string{"key": "value"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	m["new-key"] = "new value"
	fmt.Println(m)
}
