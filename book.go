package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
	price int
}

func (b Book) print() {
	fmt.Printf("%-15s: %s\n", b.title, strconv.Itoa(b.price))
}
