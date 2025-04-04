package main

import (
	"fmt"
)

func Contains[T comparable](slice []T, element T) (bool, int) {
	for i, v := range slice {
		if v == element {
			return true, i
		}
	}
	return false, -1
}

func main() {
	found, index := Contains([]int{1, 2, 3, 4}, 3)
	fmt.Printf("Found: %v at index: %d\n", found, index)

	found, index = Contains([]string{"ayan", "mehta"}, "mehta")
	fmt.Printf("Found: %v at index: %d\n", found, index)

	found, index = Contains([]string{"ayan", "mehta"}, "mehtaa")
	fmt.Printf("Found: %v at index: %d\n", found, index)
}
