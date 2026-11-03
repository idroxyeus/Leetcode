class Solution:
    def minMoves(self, nums: List[int]) -> int:
        m=max(nums)
        return (len(nums)*max(nums))-sum(nums)