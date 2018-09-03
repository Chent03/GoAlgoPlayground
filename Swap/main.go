package main

import (
	"fmt"
)

func main() {
	testArray := []int{7, 9, 4}
	fmt.Println(swap(testArray, 0, 1))

	array := []int{18, 6, 66, 44, 9, 22, 14}
	fmt.Println(indexOfMinimum(array, 2))

	sortArray := []int{22, 11, 99, 88, 9, 7, 42}
	fmt.Println(selectionSort(sortArray))
}

func swap(a []int, indexOne, indexTwo int) []int {
	a[indexOne], a[indexTwo] = a[indexTwo], a[indexOne]
	return a
}

func indexOfMinimum(a []int, startIndex int) int {
	minValue := a[startIndex]
	var minIndex int
	for i := startIndex; i < len(a); i++ {
		if minValue > a[i] {
			minValue = a[i]
			minIndex = i
		}
	}
	return minIndex
}

func selectionSort(a []int) []int {
	for i := range a {
		minIndex := indexOfMinimum(a, i)
		swap(a, i, minIndex)
	}
	return a
}
