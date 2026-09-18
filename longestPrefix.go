package main

import "fmt"

func longestCommonPrefix(strs []string) (longestPrefix string) {
	if len(strs) == 0 {
		return
	}
	shortestWord := strs[0]
	for i, v := range strs {
		if i != 0 && len(v) < len(shortestWord) {
			shortestWord = v
		}
	}
	if shortestWord == "" {
		return
	}
	for i := 0; i < len(shortestWord); i++ {
		for _, v := range strs {
			if v[i] != shortestWord[i] {
				return
			}
		}
		longestPrefix += string(shortestWord[i])
	}
	return
}

func main() {
	fmt.Println("result", longestCommonPrefix([]string{"flow", "flower", "flight"}))
}
