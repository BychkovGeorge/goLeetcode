package main

import "fmt"

func addBinary(a string, b string) (result string) {
	counter := 0
	flag := false
	for {
		if counter >= len(a) || counter >= len(b) {
			if len(a) == len(b) {
				if flag {
					result = "1" + result
				}
				return
			}
			if counter >= len(a) {
				if flag {
					flag = false
					result = addBinary(b[:len(b)-counter], "1") + result
					return
				}
				result = b[:len(b)-counter] + result
				return
			}
			if counter >= len(b) {
				if flag {
					flag = false
					result = addBinary(a[:len(a)-counter], "1") + result
					return
				}
				result = a[:len(a)-counter] + result
				return
			}
		}
		ai := a[len(a)-1-counter]
		bi := b[len(b)-1-counter]
		if ai == '0' && bi == '0' {
			if flag {
				result = "1" + result
			} else {
				result = "0" + result
			}
			flag = false
		} else if ai == '1' && bi == '1' {
			if flag {
				result = "1" + result
			} else {
				result = "0" + result
			}
			flag = true
		} else {
			if flag {
				result = "0" + result
				flag = true
			} else {
				result = "1" + result
				flag = false
			}
		}
		counter++
	}
}

func main() {
	fmt.Println("result", addBinary("11", "1"))
}
