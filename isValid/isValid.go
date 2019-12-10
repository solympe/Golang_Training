package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	stack := []uint8{}

	stack = append(stack, s[0])
	j := 1
	for i := 1; i < len(s); i++ {
		if j == 0 || s[i]-stack[j-1] > 2 || s[i] == stack[j-1] {
			stack = append(stack, s[i])
			j++
		} else {
			stack = stack[:len(stack)-1]
			j--
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false

}

func main() {

	str := "(("

	fmt.Println(isValid(str))
}
