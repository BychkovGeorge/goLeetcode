package main

import "fmt"

func main() {
	fmt.Println("result", canPlaceFlowers([]int{0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	prev := 0
	for i, v := range flowerbed {
		if v == 1 || (i > 0 && prev == 1) || (i < len(flowerbed)-1 && flowerbed[i+1] == 1) {
			prev = v
			continue
		}
		prev = 1
		n--
		if n == 0 {
			return true
		}
	}
	return false
}
