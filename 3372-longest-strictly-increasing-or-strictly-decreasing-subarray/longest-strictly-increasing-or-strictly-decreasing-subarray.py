class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        if len(nums)==0: return 0
        ans,inc,dick=1,1,1
        for R in range(1,len(nums)):
            if nums[R]>nums[R-1]:
                inc+=1
                dick=1
            elif nums[R]<nums[R-1]:
                dick+=1
                inc=1
            else:inc,dick=1,1
            ans=max(ans,inc,dick)
        return ans
        