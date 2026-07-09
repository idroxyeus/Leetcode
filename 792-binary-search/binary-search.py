class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n=len(nums)
        l=0
        m=n//2
        if nums[m]==target:
            return m
        elif nums[m]<target:
            r=n
            l=m+1
            for i in range(l,r):
                if nums[i]==target:
                    return i
                    break
        elif nums[m]>target:
            l=0
            r=m
            for i in range(l,r):
                if nums[i]==target:
                    return i
                    break
        return -1

