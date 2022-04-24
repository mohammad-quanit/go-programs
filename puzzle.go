package main

import (
	"fmt"
)

type Puzzle struct {
	title    string
	category string
}

func (p Puzzle) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.category)
}
