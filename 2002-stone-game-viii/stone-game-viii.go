func stoneGameVIII(stones []int) int {
	n := len(stones)
	prefix := make([]int, n)
	prefix[0] = stones[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	// res represents the maximum score difference the current player can get
	// starting from index i+1. Initialized with the last state prefix[n-1].
	res := prefix[n-1]
	for i := n - 2; i >= 1; i-- {
		res = max(res, prefix[i]-res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
