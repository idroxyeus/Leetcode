class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        c,ans=0,0
        for i in range(len(nums)):
            if nums[i]==1:
                c+=1
                ans=max(ans,c)
                continue
            else:
                c=0
        return ans
            