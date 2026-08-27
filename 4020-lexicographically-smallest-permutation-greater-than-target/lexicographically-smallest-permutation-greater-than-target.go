import "strings"

func lexGreaterPermutation(s string, target string) string {
	n := len(s)
	counts := make([]int, 26)
	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
	}

	// Step 1: Match the target prefix as far as possible
	L := 0
	for L < n && counts[target[L]-'a'] > 0 {
		counts[target[L]-'a']--
		L++
	}

	// If the entire target matched perfectly, we must backtrack from the last character
	if L == n {
		L--
		counts[target[L]-'a']++
	}

	// Step 2: Backtrack from right to left to find the first valid branching point
	for i := L; i >= 0; i-- {
		tChar := int(target[i] - 'a')

		// Look for the smallest available character strictly greater than target[i]
		for c := tChar + 1; c < 26; c++ {
			if counts[c] > 0 {
				var sb strings.Builder
				sb.Grow(n)
				sb.WriteString(target[:i])
				
				// Append the strictly greater character
				sb.WriteByte(byte('a' + c))
				counts[c]--

				// Fill the remaining suffix greedily with the smallest available characters
				for k := 0; k < 26; k++ {
					for counts[k] > 0 {
						sb.WriteByte(byte('a' + k))
						counts[k]--
					}
				}
				return sb.String()
			}
		}

		// If no branch was found at position i, add the previous character back to the pool
		if i > 0 {
			counts[target[i-1]-'a']++
		}
	}

	return ""
}
