class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        l,r,z,ans=0,0,0,0
        seen=set()
        for r in range(len(nums)):
            if nums[r]==0:
                z+=1
            while z>1:
                if nums[l]==0:
                    z-=1
                l+=1
            ans=max(ans,r-l)
        return ans
            
