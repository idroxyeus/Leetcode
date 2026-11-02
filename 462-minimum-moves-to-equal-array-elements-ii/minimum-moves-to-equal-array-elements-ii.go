package main

import (
	"fmt"
	"math/rand"
)

func minMoves2(nums []int) int {
	n := len(nums)
	mid := n / 2
	
	// Find the element that would be at the middle index if sorted
	median := quickselect(nums, 0, n-1, mid)
	
	moves := 0
	for _, num := range nums {
		moves += abs(num - median)
	}
	return moves
}

func quickselect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	
	// Randomize pivot to ensure average O(n) performance
	pivotIdx := left + rand.Intn(right-left+1)
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	
	if i == k {
		return nums[i]
	} else if i < k {
		return quickselect(nums, i+1, right, k)
	} else {
		return quickselect(nums, left, i-1, k)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
