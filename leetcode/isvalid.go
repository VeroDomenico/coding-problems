package main

import (
	"fmt"
)

func isValid(s string) bool {
	m := make(map[rune]rune)
	m['}'] = '{'
	m[']'] = '['
	m[')'] = '('
	//Uneven cannot be true case
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	for _, char := range s {
		if m[char] == 0 {
			stack = append(stack, rune(char))
		} else {
			if m[char] != stack[len(stack)-1] || len(stack) == 0 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Print(isValid("[]{]"))
}
