package main

import "fmt"

func main() {
	/******** Variables & conditions ***********/

	x := 10       // implicit type variable
	var y int = 4 // explicit type variable

	sum := x + y
	if sum > 9 {
		fmt.Println("is greater than 9")
	} else {
		fmt.Println("is less than 9")
	}

	var check uint32
	var bl bool
	fmt.Println(check, bl)

	fmt.Printf("%t", bl) // boolean representation/format
	fmt.Println()
	fmt.Printf("%b", 10355) // binary representation/format
	fmt.Println()
	fmt.Printf("%o", 10355) // octal representation/format
	fmt.Println()
	fmt.Printf("%d", 10355) // decimal representation/format
	fmt.Println()
	fmt.Printf("%x", 3435) // hexadecimal representation/format

	fmt.Println()
	fmt.Printf("%e", 2.5656565) // scientific notation representation/format

	fmt.Println()
	fmt.Printf("%f", 2.5656565213123213213212321) // decimal notation representation/format

	fmt.Println()
	fmt.Printf("%g", 2.5656565323123123213312) // large exponent notation representation/format

	fmt.Println()
	out := fmt.Sprintf("%07d", 45)
	fmt.Println(out)

	// /********* Array *********************/
	// a := []int{1, 2, 3, 4, 5}
	// a[2] = 33
	// a = append(a, 66)
	// fmt.Println(a)

	// var myArr = [3]string{"blue coral", "staghorn coral", "pillar coral"}
	// fmt.Println(myArr)

	// /********* Slices or dynamic array *********************/
	// sliceArr := []int{1: -3, -2, -1, 0, 1, 2, 3}
	// fmt.Println(sliceArr)

	// floatSliceArr := []float64{3.14, 9.23, 111.11, 312.12, 1}
	// floatSliceArr = append(floatSliceArr, 555)
	// fmt.Println(floatSliceArr)

	// /************ Map / Object  *****************/
	// vertices := make(map[string]int)
	// vertices["triangle"] = 3
	// vertices["square"] = 4
	// vertices["circle"] = 5
	// vertices["rectanle"] = 6
	// fmt.Println(vertices["triangle"])
	// delete(vertices, "rectanle")
	// fmt.Println(vertices)
}
