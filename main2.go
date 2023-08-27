package main

import "fmt"

func isBalance(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	s1 := "{()}{}()"
	if isBalance(s1) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}

	s2 := "{()}(])"
	if isBalance(s2) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}
