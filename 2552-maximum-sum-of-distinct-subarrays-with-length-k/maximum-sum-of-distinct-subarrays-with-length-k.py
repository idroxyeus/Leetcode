from collections import defaultdict

class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        count = defaultdict(int)
        su = 0
        for i in range(k):
            count[nums[i]] += 1
            su += nums[i]
        ma = su if len(count) == k else 0
        
        for R in range(k, len(nums)):
            count[nums[R]] += 1
            su += nums[R]
            left_val = nums[R - k]
            count[left_val] -= 1
            su -= left_val
            if count[left_val] == 0:
                del count[left_val]
            if len(count) == k:
                ma = max(ma, su)
        return ma