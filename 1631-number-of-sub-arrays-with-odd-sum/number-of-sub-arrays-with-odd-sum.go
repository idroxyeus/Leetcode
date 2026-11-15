package main

func numOfSubarrays(arr []int) int {
	const mod = 1_000_000_007

	// count[0] stores the number of even prefix sums seen so far (initialized to 1 for an empty prefix sum of 0)
	// count[1] stores the number of odd prefix sums seen so far
	count := [2]int{1, 0}
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num
		parity := prefixSum & 1 // 0 if even, 1 if odd
		result = (result + count[parity^1]) % mod
		count[parity]++
	}

	return result
}
