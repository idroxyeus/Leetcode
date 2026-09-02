class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_s,curr=nums[0],0
        for num in nums:
            if curr<0:
                curr=0
            curr+=num
            max_s=max(curr,max_s)
        return max_s