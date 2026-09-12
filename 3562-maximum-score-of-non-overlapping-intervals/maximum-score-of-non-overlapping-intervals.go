package main

import (
	"sort"
)

type Interval struct {
	start  int
	end    int
	weight int
	id     int
}

type Result struct {
	weight    int
	selection []int
}

func maximumWeight(intervals [][]int) []int {
	n := len(intervals)
	arr := make([]Interval, n)
	for i, iv := range intervals {
		arr[i] = Interval{start: iv[0], end: iv[1], weight: iv[2], id: i}
	}

	// Sort intervals by start time, and by id for tie-breaking stability
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].start != arr[j].start {
			return arr[i].start < arr[j].start
		}
		return arr[i].id < arr[j].id
	})

	// Memoization table: memo[i][count]
	memo := make([][]Result, n)
	for i := range memo {
		memo[i] = make([]Result, 5)
		for j := range memo[i] {
			memo[i][j] = Result{weight: -1}
		}
	}

	// Binary search helper to find the first non-overlapping interval
	findNext := func(currentEnd int) int {
		return sort.Search(n, func(k int) bool {
			return arr[k].start > currentEnd
		})
	}

	var solve func(int, int) Result
	solve = func(i int, count int) Result {
		if i == n || count == 0 {
			return Result{weight: 0, selection: []int{}}
		}
		if memo[i][count].weight != -1 {
			return memo[i][count]
		}

		// Option 1: Skip the current interval
		resSkip := solve(i+1, count)

		// Option 2: Take the current interval
		nextIdx := findNext(arr[i].end)
		resTake := solve(nextIdx, count-1)
		
		totalWeight := arr[i].weight + resTake.weight
		
		// Build and immediately sort the new selection to guarantee correct comparison
		currentSelection := make([]int, 0, 1+len(resTake.selection))
		currentSelection = append(currentSelection, arr[i].id)
		currentSelection = append(currentSelection, resTake.selection...)
		sort.Ints(currentSelection)

		var best Result
		if totalWeight > resSkip.weight {
			best = Result{weight: totalWeight, selection: currentSelection}
		} else if totalWeight < resSkip.weight {
			best = resSkip
		} else {
			// Weights are equal: perform the correct lexicographical tie-break
			if isLexicographicallySmaller(currentSelection, resSkip.selection) {
				best = Result{weight: totalWeight, selection: currentSelection}
			} else {
				best = resSkip
			}
		}

		memo[i][count] = best
		return best
	}

	ans := solve(0, 4)
	return ans.selection
}

func isLexicographicallySmaller(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
