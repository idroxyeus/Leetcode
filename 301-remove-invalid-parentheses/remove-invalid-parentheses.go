package main

func removeInvalidParentheses(s string) []string {
	remL, remR := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			remL++
		} else if s[i] == ')' {
			if remL > 0 {
				remL--
			} else {
				remR++
			}
		}
	}

	found := make(map[string]bool)
	var dfs func(int, int, int, int, []byte)
	dfs = func(index, rL, rR, bal int, cur []byte) {
		if bal < 0 {
			return
		}
		if index == len(s) {
			if rL == 0 && rR == 0 && bal == 0 {
				found[string(cur)] = true
			}
			return
		}

		char := s[index]
		if char == '(' && rL > 0 {
			dfs(index+1, rL-1, rR, bal, cur)
		}
		if char == ')' && rR > 0 {
			dfs(index+1, rL, rR-1, bal, cur)
		}

		cur = append(cur, char)
		if char == '(' {
			dfs(index+1, rL, rR, bal+1, cur)
		} else if char == ')' {
			dfs(index+1, rL, rR, bal-1, cur)
		} else {
			dfs(index+1, rL, rR, bal, cur)
		}
		cur = cur[:len(cur)-1]
	}

	dfs(0, remL, remR, 0, make([]byte, 0, len(s)))

	result := make([]string, 0, len(found))
	for str := range found {
		result = append(result, str)
	}
	return result
}
