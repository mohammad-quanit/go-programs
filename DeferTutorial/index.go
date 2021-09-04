package DeferTutorial

import (
	"errors"
	"fmt"
)

type Engineer struct {
	Name string
}

func DeferExample() {
	defer func() {
		fmt.Println("I'll be executed last!")
	}()

	defer func() {
		fmt.Println("last in first out!")
	}()
}

func MyAwesomeFunction() (err error) {
	defer func() {
		err = errors.New("i am an error")
	}()
	return nil
}

func ReturnValue() (x int) {
	defer func() {
		x = 10
	}()
	x = 5
	return
}

// UpdateName is the method of Engineer struct
func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

func DoStuff(e *Engineer) {
	defer e.UpdateName("Mohammad Quanit")
	fmt.Println("doing other exciting stuff!")
}
