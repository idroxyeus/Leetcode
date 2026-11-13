func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)
	var ans int64 = 0
	var cost int64 = 0
	k64 := int64(k)
	dq := []int{}
	j := n - 1

	for i := n - 1; i >= 0; i-- {
		num := nums[i]

		for len(dq) > 0 && nums[dq[len(dq)-1]] < num {
			l := dq[len(dq)-1]
			dq = dq[:len(dq)-1]
			var r int
			if len(dq) > 0 {
				r = dq[len(dq)-1]
			} else {
				r = j + 1
			}
			cost += int64(r-l) * int64(num-nums[l])
		}
		dq = append(dq, i)

		for cost > k64 {
			cost -= int64(nums[dq[0]] - nums[j])
			if dq[0] == j {
				dq = dq[1:]
			}
			j--
		}

		ans += int64(j - i + 1)
	}

	return ans
}
