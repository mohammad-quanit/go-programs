package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	// sorting array
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	fmt.Println(arr)

	left := 0
	right := len(arr) - 1

	if right < left {
		fmt.Println("Not found")
		return -1
	}

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			fmt.Println("Found at position: ", mid)
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
