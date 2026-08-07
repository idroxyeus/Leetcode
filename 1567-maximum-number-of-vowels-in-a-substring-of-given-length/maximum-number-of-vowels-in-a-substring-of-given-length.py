class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        ma=0
        v={'a','e','i','o','u'}
        for i in range(k):
            if s[i] in v:
                ma+=1
        mv=ma
        for R in range(k,len(s)):
            if s[R] in v:
                ma+=1
            if s[R-k] in v:
                ma-=1
            mv=max(mv,ma)
        return mv