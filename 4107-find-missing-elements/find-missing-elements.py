class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        ma,mi,se=max(nums),min(nums),set(nums)
        ans=[]
        for num in range(mi,ma+1):
            if num not in se:
                ans.append(num)
        return ans


            