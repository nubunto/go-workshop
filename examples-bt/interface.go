package main

import "fmt"

func main() {
	var myInt interface{} = 1
	fmt.Println(myInt.(string))
}
