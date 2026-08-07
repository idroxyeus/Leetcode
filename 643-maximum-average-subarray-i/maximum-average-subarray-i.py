class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        su=sum(nums[:k])
        ma=su
        for R in range(k,len(nums)):
            su=su+nums[R]-nums[R-k]
            ma=max(su,ma)
        return ma/k