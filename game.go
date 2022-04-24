package main

import (
	"fmt"
	"strconv"
)

type Game struct {
	title string
	price int
}

func (g Game) print() {
	fmt.Printf("%-15s: %s\n", g.title, strconv.Itoa(g.price))
}
