class Solution:
    def maxProduct(self, n: int) -> int:
        res=0
        l=[]
        while n>0:
            l.append(n%10)
            n=n//10
        l.sort()
        return l[-1]*l[-2]
