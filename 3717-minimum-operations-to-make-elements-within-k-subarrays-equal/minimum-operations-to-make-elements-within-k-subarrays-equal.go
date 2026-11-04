package main

import (
	"container/heap"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}){ *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// A MaxHeap is a max-heap of ints.
type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}){ *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)
	if n < x*k {
		return -1 // Not enough elements to pick k non-overlapping windows of size x
	}

	minOps := make([]int64, n-x+1)
	
	leftHeap := &MaxHeap{}  // Stores the smaller half
	rightHeap := &IntHeap{} // Stores the larger half
	heap.Init(leftHeap)
	heap.Init(rightHeap)
	
	var leftSum, rightSum int64
	
	// Lazy deletion maps
	leftDel := make(map[int]int)
	rightDel := make(map[int]int)
	
	var leftSize, rightSize int

	// Helper to prune invalid elements from the top of the heaps
	prune := func() {
		for leftHeap.Len() > 0 && leftDel[(*leftHeap)[0]] > 0 {
			leftDel[(*leftHeap)[0]]--
			heap.Pop(leftHeap)
		}
		for rightHeap.Len() > 0 && rightDel[(*rightHeap)[0]] > 0 {
			rightDel[(*rightHeap)[0]]--
			heap.Pop(rightHeap)
		}
	}

	// Helper to maintain heap sizes
	balance := func() {
		targetLeft := (x + 1) / 2
		prune()
		
		for leftSize < targetLeft && rightSize > 0 {
			val := heap.Pop(rightHeap).(int)
			rightSum -= int64(val)
			rightSize--
			
			heap.Push(leftHeap, val)
			leftSum += int64(val)
			leftSize++
			prune()
		}
		for leftSize > targetLeft {
			val := heap.Pop(leftHeap).(int)
			leftSum -= int64(val)
			leftSize--
			
			heap.Push(rightHeap, val)
			rightSum += int64(val)
			rightSize++
			prune()
		}
	}
	
	add := func(val int) {
		if leftHeap.Len() == 0 || val <= (*leftHeap)[0] {
			heap.Push(leftHeap, val)
			leftSum += int64(val)
			leftSize++
		} else {
			heap.Push(rightHeap, val)
			rightSum += int64(val)
			rightSize++
		}
		balance()
	}
	
	remove := func(val int) {
		if val <= (*leftHeap)[0] {
			leftDel[val]++
			leftSum -= int64(val)
			leftSize--
		} else {
			rightDel[val]++
			rightSum -= int64(val)
			rightSize--
		}
		balance()
	}
	
	// Process the initial window
	for i := 0; i < x; i++ {
		add(nums[i])
	}
	
	// Calculate cost for all windows
	for i := 0; i <= n-x; i++ {
		prune()
		median := int64((*leftHeap)[0])
		
		lCount := int64((x + 1) / 2)
		rCount := int64(x - int(lCount))
		
		cost := (median*lCount - leftSum) + (rightSum - median*rCount)
		minOps[i] = cost
		
		if i+x < n {
			remove(nums[i])
			add(nums[i+x])
		}
	}
	
	// DP Step
	const INF = math.MaxInt64 / 2
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	
	// Base Cases
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	
	for j := 1; j <= k; j++ {
		for i := n - x; i >= 0; i-- {
			// Choice 1: Skip current element
			dp[i][j] = dp[i+1][j]
			
			// Choice 2: Form a window starting at index i
			if dp[i+x][j-1] != INF {
				cost := minOps[i] + dp[i+x][j-1]
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}
	
	if dp[0][k] >= INF {
		return -1
	}
	return dp[0][k]
}
