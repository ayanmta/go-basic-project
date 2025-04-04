package main

import (
	"fmt"
	"strings"
)

func Contains[T comparable](slice []T, element T, caseSensitive bool) bool {
	if str, ok := any(element).(string); ok && !caseSensitive {
		for _, v := range slice {
			if vStr, ok := any(v).(string); ok {
				if strings.EqualFold(vStr, str) {
					return true
				}
			}
		}
		return false
	}

	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3, 4}, 3, true))
	fmt.Println(Contains([]string{"ayan", "mehta"}, "MEHTA", false))
	fmt.Println(Contains([]string{"ayan", "mehta"}, "mehtaa", true))
}
