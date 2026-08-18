func largestInteger(nums []int, k int) int {
    n := len(nums)
    if k == 1 {
        cnt := make(map[int]int)
        for _, x := range nums {
            cnt[x]++
        }
        ans := -1
        for x, v := range cnt {
            if v == 1 && x > ans {
                ans = x
            }
        }
        return ans
    }
    if k == n {
        maxVal := nums[0]
        for _, x := range nums {
            if x > maxVal {
                maxVal = x
            }
        }
        return maxVal
    }
    
    check := func(idx int) int {
        val := nums[idx]
        for i, x := range nums {
            if i != idx && x == val {
                return -1
            }
        }
        return val
    }
    
    ans := check(0)
    if last := check(n - 1); last > ans {
        ans = last
    }
    return ans
}
