package main

import (
	"fmt"
)

func main() {
	target := 6
	arr := []int{2, 4, 7, 8, 3, 3, 5, 6}

	// Linear Search example
	index := BinarySearch(arr, target)
	if index == -1 {
		fmt.Printf("No target %d found in provided array", target)
	} else {
		fmt.Printf("Target %d found in index %d\n", target, index)
	}

}
