class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        L,R=0,0
        seen=set()
        for R in range(len(nums)):
            if nums[R] in seen:
                return True
            seen.add(nums[R])
            if len(seen)>k:
                seen.remove(nums[L])
                L+=1
        return False