package main

import (
	"sort"
)

type pair struct {
	val int
	idx int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	pairs := make([]pair, n)
	for i, v := range nums {
		pairs[i] = pair{val: v, idx: i}
	}

	// 1. Sort pairs based on their values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	result := make([]int, n)
	
	// 2. Group connected components where adjacent elements differ by <= limit
	i := 0
	for i < n {
		j := i + 1
		for j < n && pairs[j].val-pairs[j-1].val <= limit {
			j++
		}

		// Extract original indices for the current connected group
		indices := make([]int, j-i)
		for k := i; k < j; k++ {
			indices[k-i] = pairs[k].idx
		}

		// Sort original indices to place smaller values at smaller index positions
		sort.Ints(indices)

		// 3. Assign sorted values to sorted indices greedily
		for k := i; k < j; k++ {
			result[indices[k-i]] = pairs[k].val
		}

		i = j // Move to the next component
	}

	return result
}
