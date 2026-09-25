package main

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	uniqueWords := make(map[string]struct{})

	var dfs func(exp string)
	dfs = func(exp string) {
		j := strings.Index(exp, "}")
		if j == -1 {
			uniqueWords[exp] = struct{}{}
			return
		}
		
		// Find the matching opening brace '{' before index j
		i := strings.LastIndex(exp[:j], "{")
		
		a := exp[:i]
		c := exp[j+1:]
		
		// Split comma-separated options inside the braces {b1,b2,...}
		options := strings.Split(exp[i+1:j], ",")
		for _, b := range options {
			dfs(a + b + c)
		}
	}

	dfs(expression)

	// Extract keys from map and sort them alphabetically
	var result []string
	for word := range uniqueWords {
		result = append(result, word)
	}
	sort.Strings(result)

	return result
}
