class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res,s=[],[]
        nums.sort()
        def sub(i):
            if i>=len(nums):
                res.append(s.copy())
                return
            s.append(nums[i])
            sub(i+1)
            s.pop()
            while i+1<len(nums) and nums[i+1]==nums[i]:
                i+=1
            sub(i+1)
        sub(0)
        return res