package main

import (
	"fmt"
)

func main() {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
		41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	fmt.Println(doSearch(primes, 73))
}

func doSearch(a []int, target int) bool {
	min := 0
	max := len(a) - 1
	var guess int
	for min <= max {
		guess = (max + min) / 2
		if a[guess] == target {
			return true
		} else if a[guess] < target {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}
	return false
}
