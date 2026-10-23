package main

func sortPermutation(nums []int) int {
    // Initialize ans to -1 (in two's complement, all bits are set to 1)
    // This serves as the identity element for the bitwise AND operation.
    ans := -1
    
    for i, x := range nums {
        if i != x {
            ans &= x
        }
    }
    
    // If ans remains -1, it means no element was misplaced (array already sorted).
    // In that case, or if the resulting ans is negative, we return 0.
    if ans == -1 {
        return 0
    }
    
    return ans
}
