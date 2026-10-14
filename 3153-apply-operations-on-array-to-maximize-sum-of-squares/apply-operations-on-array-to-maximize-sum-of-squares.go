package main

func maxSum(nums []int, k int) int {
    const MOD = 1_000_000_007
    // Array to store the count of set bits at each position (up to 30 bits for nums[i] <= 10^9)
    bitCounts := make([]int, 32)
    
    // Count the total occurrences of each bit position across all numbers
    for _, num := range nums {
        for i := 0; i < 30; i++ {
            if (num >> i) & 1 == 1 {
                bitCounts[i]++
            }
        }
    }
    
    ans := 0
    
    // Greedily reconstruct the k largest possible numbers
    for idx := 0; idx < k; idx++ {
        var current int64 = 0
        for i := 0; i < 30; i++ {
            if bitCounts[i] > 0 {
                current |= (1 << i)
                bitCounts[i]--
            }
        }
        // Calculate square and add to final answer modulo 10^9 + 7
        square := (current * current) % MOD
        ans = (ans + int(square)) % MOD
    }
    
    return ans
}