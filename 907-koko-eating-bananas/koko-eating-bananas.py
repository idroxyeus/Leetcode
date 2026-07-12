class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        max_k=max(piles)
        if len(piles)==h:
            return max_k
        l,r=1,max_k
        while l<=r:
            m=(l+r)//2
            hours_needed=0
            for p in piles:
                hours_needed+= math.ceil(float(p) / m)
            if hours_needed>h:
                l=m+1
            elif hours_needed<=h:
                r=m-1
        return l
