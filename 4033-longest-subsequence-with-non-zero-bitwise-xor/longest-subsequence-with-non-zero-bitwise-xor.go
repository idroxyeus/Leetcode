package main

func longestSubsequence(nums []int) int {
	totalXOR := 0
	hasNonZero := false

	for _, num := range nums {
		totalXOR ^= num
		if num != 0 {
			hasNonZero = true
		}
	}

	if totalXOR != 0 {
		return len(nums)
	}
	if !hasNonZero {
		return 0
	}
	return len(nums) - 1
}
