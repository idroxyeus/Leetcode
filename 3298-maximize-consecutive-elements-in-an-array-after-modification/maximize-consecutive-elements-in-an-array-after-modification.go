package main

import "slices"

func maxSelectedElements(nums []int) int {
	slices.Sort(nums)
	dp := make(map[int]int)
	ans := 0

	for _, num := range nums {
		dp[num+1] = dp[num] + 1
		dp[num] = dp[num-1] + 1
		ans = max(ans, dp[num], dp[num+1])
	}

	return ans
}
