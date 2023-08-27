package main

import "fmt"

func isBalanced(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			lastOpen := stack[len(stack)-1]

			if (char == ')' && lastOpen == '(') ||
				(char == '}' && lastOpen == '{') ||
				(char == ']' && lastOpen == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	// Вводим строку со скобками
	input := "{()}{}()"

	if isBalanced(input) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}
