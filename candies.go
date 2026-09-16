package main

import "fmt"

func main() {
	fmt.Println("result", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxVal int
	result := make([]bool, len(candies))

	for _, v := range candies {
		if v > maxVal {
			maxVal = v
		}
	}

	for i, v := range candies {
		result[i] = v+extraCandies >= maxVal
	}

	return result
}
