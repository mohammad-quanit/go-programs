package main

import (
	"fmt"

	"github.com/mohammad-quanit/go-programs/exercises"
)

func main() {
	str1 := "hello"
	str2 := "world"
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)

	fmt.Println()
	exercises.StringToNumeric()

	fmt.Println()
	exercises.NumericToString()

	fmt.Println()
	exercises.StringParser()

	fmt.Println()
	exercises.StringLength("welcome to go")

	fmt.Println()
	exercises.StringCopy()

}
