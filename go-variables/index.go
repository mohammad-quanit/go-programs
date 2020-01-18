package main

import "fmt"

func main()  {
	/******** Variables & conditions ***********/
	x := 5
	var y int = 5
	sum := x + y
	if sum > 9 {
		fmt.Println("is greater than 9")
	} else {
		fmt.Println("is less than 9")
	}

	/********* Array *********************/
	a := []int{1, 2, 3, 4, 5}
	a[2] = 33
	a = append(a, 66)
	fmt.Println(a)

	/************ Map / Object  *****************/
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["circle"] = 5
	vertices["rectanle"] = 6
	fmt.Println(vertices["triangle"])
	delete(vertices, "rectanle")
	fmt.Println(vertices)
}