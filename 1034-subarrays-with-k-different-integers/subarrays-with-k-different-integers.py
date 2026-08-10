class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        def norm(m):
            seen=dict()
            l,r,ans=0,0,0
            for r in range(len(nums)):
                if nums[r] in seen:
                    seen[nums[r]]+=1
                else:
                    seen[nums[r]]=1
                while len(seen)>m:
                    if nums[l] in seen:
                        seen[nums[l]]-=1
                    if seen[nums[l]]==0:
                        del seen[nums[l]]
                    l+=1
                ans+=r-l+1
            return ans
        return norm(k)-norm(k-1)