package main

import (
	// "fmt"
	"fmt"

	"github.com/mohammad-quanit/go-programs/DeferTutorial"
)

func main() {
	fmt.Println("I'll be executed first!")

	// DeferTutorial.DeferExample()

	// err := DeferTutorial.MyAwesomeFunction()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(DeferTutorial.ReturnValue())

	quanit := &DeferTutorial.Engineer{Name: "Mohammad"}
	fmt.Printf("%+v\n", quanit)
	DeferTutorial.DoStuff(quanit)
	fmt.Printf("%+v\n", quanit)
}
