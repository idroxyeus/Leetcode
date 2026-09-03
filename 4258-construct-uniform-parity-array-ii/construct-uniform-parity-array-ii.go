package main

func uniformArray(nums1 []int) bool {
    mn := int(^uint(0) >> 1) // MaxInt
    for _, x := range nums1 {
        if x%2 == 1 && x < mn {
            mn = x
        }
    }
    
    // If no odd number exists, all elements are already even.
    if mn == int(^uint(0)>>1) {
        return true
    }
    
    // Check if any even number is smaller than the minimum odd number
    for _, x := range nums1 {
        if x%2 == 0 && x < mn {
            return false
        }
    }
    
    return true
}