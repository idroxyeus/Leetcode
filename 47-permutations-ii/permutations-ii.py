class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res,digit,chosen=[],[],[False]*len(nums)
        def perm():
            if len(digit)==len(nums):
                res.append(digit.copy())
                return
            for i in range(len(nums)):
                if chosen[i]: continue
                if i>0 and nums[i]==nums[i-1] and not chosen[i-1]:
                    continue
                chosen[i]=True
                digit.append(nums[i])
                perm()
                digit.pop()
                chosen[i]=False
        perm()
        return res