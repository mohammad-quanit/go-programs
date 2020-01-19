package main

import (
	"fmt"
	"math"
	"math/rand"
)

/********* Math Object *********************/
func mathFunc() {
	fmt.Println("The square root is", math.Sqrt(4))
	fmt.Println("The random number from (1-100) is", rand.Intn(100))
}

// ******* return type functions add ************/
func add(x, y float64) float64 {
	return x * y
}

// ******* return type functions multiple values ************/
func mutipleVals(x, y string) (string, string) {
	return x, y
}

func main() {
	var π float32 = 3.14
	fmt.Println(5 + π)

	var maxUint32 uint32 = 4294967295  // Max uint32 size
	maxUint32 = maxUint32 + 1
  fmt.Println(maxUint32)

	// mathFunc()
	// // var num1, num2 float64 = 5.6, 100.55
	// num1, num2 := 5.6, 1000.66
	// fmt.Println(add(num1, num2))

	// w1, w2 := "Hey", "There"
	// fmt.Println(mutipleVals(w1, w2))

	// // explicitly type inference
	// var isInt = 65
	// var b float64 = float64(isInt)
	// fmt.Println(b)

	// // implicitly type inference
	// isInt2 := 55
	// isFloat64 := isInt2
	// fmt.Println(isFloat64)
}
