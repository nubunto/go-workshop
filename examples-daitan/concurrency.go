package main

import "fmt"

func printMessage(ch chan int, message string) {
	fmt.Println(message)
	ch <- 1
}

func main() {
	ch := make(chan int)
	go printMessage(ch, "e ai blz")
	<-ch
}
