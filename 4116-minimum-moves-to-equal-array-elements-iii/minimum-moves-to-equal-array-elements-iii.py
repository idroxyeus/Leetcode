class Solution:
    def minMoves(self, nums: List[int]) -> int:
        m=max(nums)
        count=0
        for i in range(len(nums)):
            count+=(m-nums[i])
        return count