package main

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	k := k1 + k2
	const maxDiff = 100000
	bucket := make([]int, maxDiff+1)
	
	var totalDiff int64 = 0
	for i := 0; i < len(nums1); i++ {
		diff := abs(nums1[i] - nums2[i])
		if diff > 0 {
			bucket[diff]++
			totalDiff += int64(diff)
		}
	}
	
	if totalDiff <= int64(k) {
		return 0
	}
	
	for d := maxDiff; d > 0; d-- {
		if bucket[d] == 0 {
			continue
		}
		
		count := bucket[d]
		if k >= count {
			k -= count
			bucket[d-1] += count
			bucket[d] = 0
		} else {
			bucket[d-1] += k
			bucket[d] -= k
			k = 0
			break
		}
	}
	
	var ans int64 = 0
	for d := 1; d <= maxDiff; d++ {
		if bucket[d] > 0 {
			ans += int64(bucket[d]) * int64(d) * int64(d)
		}
	}
	
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
