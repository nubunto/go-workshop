package main

import "fmt"

type ageSayer interface {
	sayMyAge()
}

type human struct {
	name string
	age  int
}

type employee struct {
	human
	companyName string
}

func (h human) sayMyAge() {
	fmt.Println("my age is", h.age)
}

func main() {
	e := employee{
		human: human{
			age: 18,
		},
		companyName: "BT",
	}
	var a ageSayer = e
	a.sayMyAge()
}
