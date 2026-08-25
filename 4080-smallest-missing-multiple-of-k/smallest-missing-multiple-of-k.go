// Do NOT include "package main" or "func main()" when pasting into LeetCode.

func missingMultiple(nums []int, k int) int {
    numSet := make(map[int]bool, len(nums))
    for _, num := range nums {
        numSet[num] = true
    }

    ans := k
    // In Go, map lookups return the value (true) or default (false) if missing.
    for numSet[ans] {
        ans += k
    }

    return ans
}
