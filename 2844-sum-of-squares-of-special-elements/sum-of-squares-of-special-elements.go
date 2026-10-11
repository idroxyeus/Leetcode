func sumOfSquares(nums []int) int {
    n := len(nums)
    sum := 0
    for i := 0; i < n; i++ {
        if n%(i+1) == 0 {
            sum += nums[i] * nums[i]
        }
    }
    return sum
}