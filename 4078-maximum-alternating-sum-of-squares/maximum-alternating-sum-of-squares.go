func maxAlternatingSum(nums []int) int64 {
    for i, x := range nums {
		nums[i] = x * x
	}

	slices.Sort(nums)

	var ans int64 = 0
	n := len(nums)
	m := n / 2

	for _, x := range nums[:m] {
		ans -= int64(x)
	}

	for _, x := range nums[m:] {
		ans += int64(x)
	}

	return ans
}