package main

import (
	"fmt"
)

func isValid(s string) bool {
	result := make([]byte, 0, len(s))
	oppositeMap := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			result = append(result, s[i])
		case ')', '}', ']':
			if len(result) == 0 || result[len(result)-1] != oppositeMap[s[i]] {
				return false
			}
			result = result[:len(result)-1]
		}
	}

	return len(result) == 0
}

func main() {
	fmt.Println("result", isValid("{([)]}"))
}
