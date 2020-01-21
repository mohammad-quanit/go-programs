package main

import "fmt"

func main() {
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

	var myArr = [3]string{"blue coral", "staghorn coral", "pillar coral"}
	fmt.Println(myArr)

	/********* Slices or dynamic array *********************/
	sliceArr := []int{1: -3, -2, -1, 0, 1, 2, 3}
	fmt.Println(sliceArr)

	floatSliceArr := []float64{3.14, 9.23, 111.11, 312.12, 1}
	floatSliceArr = append(floatSliceArr, 555)
	fmt.Println(floatSliceArr)

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
