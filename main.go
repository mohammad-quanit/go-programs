package main

import (
	"fmt"
	"sort"
)

func isHandOfStraights(hand []int, k int) bool {
	if len(hand)%k != 0 {
		return false
	}
	count := make(map[int]int)
	for _, v := range hand {
		count[v] = count[v] + 1
	}
	sort.Ints(hand)
	i := 0
	n := len(hand)
	for i < n {
		current := hand[i]
		for j := 0; j < k; j++ {
			if _, ok := count[current+j]; !ok || count[current+j] == 0 {
				return false
			}
			count[current+j]--
		}
		for i < n && count[hand[i]] == 0 {
			i++
		}
	}
	return true
}

func main() {
	hand := []int{5, 2, 4, 4, 1, 3, 5, 6, 3}
	k := 3
	fmt.Println(isHandOfStraights(hand, k))
}
