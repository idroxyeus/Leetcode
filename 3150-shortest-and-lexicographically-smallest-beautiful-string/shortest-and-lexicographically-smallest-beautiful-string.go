package main

import (
	"fmt"
)

func shortestBeautifulSubstring(s string, k int) string {
	ans := ""
	left := 0
	onesCount := 0

	// Right pointer expands the window
	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			onesCount++
		}

		// Shrink the window from the left if we have exactly k ones
		// or if we have exceeded k ones.
		for onesCount == k {
			// A valid beautiful substring must start and end with '1' to be the shortest possible
			if s[left] == '1' {
				currentSub := s[left : right+1]
				
				// Update answer if it's the first valid substring found,
				// if it's shorter, or if it's the same length but lexicographically smaller.
				if ans == "" || len(currentSub) < len(ans) || (len(currentSub) == len(ans) && currentSub < ans) {
					ans = currentSub
				}
				
				// Move left pointer and reduce the count to break the loop condition
				onesCount--
				left++
				break
			}
			left++
		}
	}

	return ans
}

