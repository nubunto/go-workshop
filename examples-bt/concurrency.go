package main

import "fmt"

func goSay(ch chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "say this!"
	}()
	goSay(ch)
}
