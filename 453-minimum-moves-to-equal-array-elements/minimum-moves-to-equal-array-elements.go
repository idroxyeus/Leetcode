package main

import "math"

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minVal := math.MaxInt64
	sum := 0

	for _, num := range nums {
		sum += num
		if num < minVal {
			minVal = num
		}
	}

	return sum - minVal*len(nums)
}
