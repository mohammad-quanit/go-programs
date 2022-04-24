package main

import (
	"fmt"
	"strings"
)

type Device interface {
	turnOn() string
	update(version float32)
}

type IPhone struct {
	name    string
	model   string
	version float32
}

type IMac struct {
	name    string
	model   string
	version float32
}

func (iphone IPhone) turnOn() string {
	return "iOS starting up..."
}

func (iphone *IPhone) update(version float32) {
	iphone.version = version
}

func (imac IMac) turnOn() string {
	return "iMac starting up..."
}

func (imac *IMac) update(version float32) {
	imac.version = version
}

func main() {
	fmt.Println("Learning Interfaces")

	// ************** Example 1 ***************
	device1 := IPhone{"IPHONE", "x pro", 13.1}
	device2 := IMac{"IMAC", "larger t3", 10.15}

	devices := []Device{&device1, &device2}

	for _, d := range devices {
		if strings.Contains(d.turnOn(), "iOS") {
			d.update(14.0)
		} else if strings.Contains(d.turnOn(), "iMac") {
			d.update(11.00)
		}
	}
	fmt.Println(device1)
	fmt.Println(device2)

	// ************** Example 2 ***************

	// var (
	// 	game   = Game{title: "minecraft", price: 10}
	// 	book   = Book{title: "Rich Dad poor Dad", price: 20}
	// 	puzzle = Puzzle{title: "Amazing Puzzle", category: "Adventure"}
	// )

	// var store List

	// store = append(store, game, book, puzzle)
	// store.print()

	// // interface values are comparable
	// fmt.Println(store[0] == game)
}
