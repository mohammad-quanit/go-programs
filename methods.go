package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y int
	r    float64
}

func (c Circle) display() {
	fmt.Printf("x=%d,y=%d,r=%.2f\n", c.x, c.y, c.r)
}

func (c Circle) area() {
	fmt.Println(math.Pi * math.Pow(c.r, 2))
}

func (c Circle) moveTo(newX, newY int) {
	c.x = newX
	c.y = newY
	fmt.Println("After moving", c.x, c.y)
}
