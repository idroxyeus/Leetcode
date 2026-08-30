package main

func minimumDeletions(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }

    // Step 1: Find the indices of the min and max elements
    minIdx, maxIdx := 0, 0
    for i := 1; i < n; i++ {
        if nums[i] < nums[minIdx] {
            minIdx = i
        }
        if nums[i] > nums[maxIdx] {
            maxIdx = i
        }
    }

    // Step 2: Ensure left is the smaller index and right is the larger index
    left, right := minIdx, maxIdx
    if left > right {
        left, right = right, left
    }

    // Step 3: Calculate deletions for the three strategies
    // Strategy 1: Delete both from the front
    delFront := right + 1

    // Strategy 2: Delete both from the back
    delBack := n - left

    // Strategy 3: Delete one from front and one from back
    delBoth := (left + 1) + (n - right)

    // Return the minimum of the three strategies
    return min(delFront, min(delBack, delBoth))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
