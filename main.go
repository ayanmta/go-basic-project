package main

import (
	"fmt"
	"strings"
)

func Contains[T comparable](slice []T, element T, caseSensitive bool) (bool, int) {
	if str, ok := any(element).(string); ok && !caseSensitive {
		for i, v := range slice {
			if vStr, ok := any(v).(string); ok {
				if strings.EqualFold(vStr, str) {
					return true, i
				}
			}
		}
		return false, -1
	}

	for i, v := range slice {
		if v == element {
			return true, i
		}
	}
	return false, -1
}

func main() {
	found, index := Contains([]int{1, 2, 3, 4}, 3, true)
	fmt.Printf("Found: %v at index: %d\n", found, index)

	found, index = Contains([]string{"ayan", "mehta"}, "MEHTA", false)
	fmt.Printf("Found: %v at index: %d\n", found, index)

	found, index = Contains([]string{"ayan", "mehta"}, "mehtaa", true)
	fmt.Printf("Found: %v at index: %d\n", found, index)
}
