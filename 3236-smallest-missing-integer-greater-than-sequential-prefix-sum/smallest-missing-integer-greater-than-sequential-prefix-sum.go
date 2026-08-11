func missingInteger(nums []int) int {
    s := nums[0]
    for j := 1; j < len(nums) && nums[j] == nums[j-1]+1; j++ {
        s += nums[j]
    }
    
    vis := map[int]bool{}
    for _, x := range nums {
        vis[x] = true
    }
    
    for x := s; ; x++ {
        if !vis[x] {
            return x
        }
    }
}