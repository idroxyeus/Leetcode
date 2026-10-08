package main

import (
	"strings"
)

func removeOuterParentheses(s string) string {
	var result strings.Builder
	opened := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			// If it's already nested, include it
			if opened > 0 {
				result.WriteByte(char)
			}
			opened++
		} else if char == ')' {
			opened--
			// If it's still nested, include it
			if opened > 0 {
				result.WriteByte(char)
			}
		}
	}

	return result.String()
}
