package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalanced(input string) string {
	var stack []rune

	for _, char := range input {
		if char == ' ' {
			continue
		}

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return "NO"
			}

			top := stack[len(stack)-1]

			if (char == ')' && top != '(') ||
				(char == ']' && top != '[') ||
				(char == '}' && top != '{') {
				return "NO"
			}

			stack = stack[:len(stack)-1]
		} else {
			return "NO"
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Print("Masukkan bracket: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	result := isBalanced(input)
	fmt.Println("Output:", result)
}
