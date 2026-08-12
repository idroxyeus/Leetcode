class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        seen=dict()
        l,ans,m=0,0,0
        for r in range(len(nums)):
            if nums[r] in seen:
                seen[nums[r]]+=1
            else:
                seen[nums[r]]=1
            while seen[nums[r]]>k:
                if nums[l] in seen:
                    seen[nums[l]]-=1
                else:
                    del seen[nums[l]]
                l+=1
            ans=max(ans,r-l+1)
        return ans