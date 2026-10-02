package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	left := 0
	right := int(math.Floor(float64(x / 2)))

	for {
		if right-left <= 1 {
			return right
		}

		middle := int(math.Floor(float64((right + left) / 2)))
		square := middle * middle
		lowSquare := (middle - 1) * (middle - 1)
		highSquare := (middle + 1) * (middle + 1)
		if square == x {
			return middle
		}
		if square < x {
			if highSquare > x {
				return middle
			}
			if highSquare == x {
				return middle + 1
			}
			left = middle
		} else {
			if lowSquare <= x {
				return middle - 1
			}
			right = middle
		}
	}
}

func main() {
	fmt.Println("result", mySqrt(5))
}
