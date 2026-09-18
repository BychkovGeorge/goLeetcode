package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}

func main() {
	fmt.Println("result", plusOne([]int{9, 8, 9}))
}
