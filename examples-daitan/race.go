package main

import (
	"fmt"
	"sync"
)

var myInt int

func main() {
	myInt = 10

	m := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			m.Lock()
			myInt++
			m.Unlock()
		}()
	}

	m.Lock()
	fmt.Println(myInt)
	m.Unlock()
}
