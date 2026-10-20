package main

func isSubarraySortable(sub []int) (int, int, bool) {
    n := len(sub)
    if n == 1 {
        return sub[0], sub[0], true
    }
    
    // Find drop points in circular fashion
    drops := 0
    for i := 0; i < n; i++ {
        if sub[i] > sub[(i+1)%n] {
            drops++
        }
    }
    if drops > 1 {
        return 0, 0, false
    }

    // Find min and max
    minVal, maxVal := sub[0], sub[0]
    for _, v := range sub {
        if v < minVal {
            minVal = v
        }
        if v > maxVal {
            maxVal = v
        }
    }
    return minVal, maxVal, true
}

func checkK(nums []int, k int) bool {
    n := len(nums)
    var prevMax *int

    for i := 0; i < n; i += k {
        sub := nums[i : i+k]
        minVal, maxVal, ok := isSubarraySortable(sub)
        if !ok {
            return false
        }
        if prevMax != nil && minVal < *prevMax {
            return false
        }
        prevMax = &maxVal
    }
    return true
}

func sortableIntegers(nums []int) int {
    n := len(nums)
    totalSum := 0

    for k := 1; k <= n; k++ {
        if n%k == 0 {
            if checkK(nums, k) {
                totalSum += k
            }
        }
    }

    return totalSum
}
