package main

import (
	"fmt"
)

func main() {
	x := 15
	a := x  // Here is the value
	b := &x // here is the memory address
	fmt.Println(a)
	fmt.Println(b)
	*b = 5
	fmt.Println(x, *b)
	*b = *b * *b
	fmt.Println(x, *b)
}
