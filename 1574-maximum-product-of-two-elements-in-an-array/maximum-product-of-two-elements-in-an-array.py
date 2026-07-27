class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        m1,m2=0,0
        for num in nums:
            if num>m2:
                m1=m2
                m2=num
            elif num>m1:
                m1=num
        return (m1-1)*(m2-1)