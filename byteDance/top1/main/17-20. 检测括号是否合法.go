/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func good(a, b uint8) bool {
	return a == '(' && b == ')' ||
		a == '[' && b == ']' ||
		a == '{' && b == '}'
}

func isValid(s string) bool {
	var stack = []uint8{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) <= 0 {
				return false
			}
			var char = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !good(char, s[i]) {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[[["))
}
