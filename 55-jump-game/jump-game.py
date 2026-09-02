class Solution:
    def canJump(self, nums: List[int]) -> bool:
        goal=len(nums)-1
        i=goal-1
        while i>=0:
            if nums[i]+i>=goal:
                goal=i
            i-=1
        return goal==0