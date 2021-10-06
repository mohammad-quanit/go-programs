package main

import (
	"time"
)

type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func main() {
	var emp Employee
	newEmp := new(Employee)

	// set values
	emp.id = 2
	emp.name = "Employee 1"
	emp.country = "DE"
	emp.created = time.Now()

	newEmp.id = 5
	newEmp.name = "Employee 2"
	newEmp.country = "BE"
	newEmp.created = time.Now()

	shape := Circle{x: 12, y: 15, r: 22.5}
	shape.display()
	shape.area()
	shape.moveTo(66, 77)
}
