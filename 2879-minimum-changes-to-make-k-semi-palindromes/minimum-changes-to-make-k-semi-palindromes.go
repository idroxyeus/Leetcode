package main

import "math"

func minimumChanges(s string, k int) int {
	n := len(s)
	inf := math.MaxInt32 / 2

	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = inf
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			length := j - i + 1
			for d := 1; d*d <= length; d++ {
				if length%d == 0 {
					c1 := getChanges(s, i, j, d)
					if c1 < cost[i][j] {
						cost[i][j] = c1
					}
					d2 := length / d
					if d2 < length && d2 != d {
						c2 := getChanges(s, i, j, d2)
						if c2 < cost[i][j] {
							cost[i][j] = c2
						}
					}
				}
			}
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for p := 0; p < i-1; p++ {
				if dp[p][j-1] != inf && cost[p][i-1] != inf {
					if dp[p][j-1]+cost[p][i-1] < dp[i][j] {
						dp[i][j] = dp[p][j-1] + cost[p][i-1]
					}
				}
			}
		}
	}

	return dp[n][k]
}

func getChanges(s string, l, r, d int) int {
	changes := 0
	for rem := 0; rem < d; rem++ {
		left := l + rem
		right := r - (d - 1 - rem)
		for left < right {
			if s[left] != s[right] {
				changes++
			}
			left += d
			right -= d
		}
	}
	return changes
}
