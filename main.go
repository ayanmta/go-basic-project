package main

import (
	"fmt"
)

func Contains[T comparable](slice []T, element T) bool {

	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3, 4}, 3))
	fmt.Println(Contains([]string{"ayan", "mehta"}, "mehtaa"))
}
