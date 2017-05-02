package main

import "fmt"

// START OMIT
// let's define a new type which is a struct
type ThisStruct struct {
	anInteger int
}

// this function returns a pointer to a ThisStruct type
func NewThisStruct() *ThisStruct {
	// the & operator captures the address of the variable (who remembers C? ;)
	// it's safe to return a pointer here, since the runtime takes care
	// of the memory allocation
	return &ThisStruct{
		anInteger: 10,
	}
}

func main() {
	ts := NewThisStruct()
	fmt.Println("an integer:", ts.anInteger)
}

// END OMIT
