from collections import Counter
class Solution:
    def maxConsecutiveAnswers(self, answerKey: str, k: int) -> int:
        
        def check(s):
            l,r,t,ans_T=0,0,0,0
            for r in range(len(answerKey)):
                if answerKey[r]==s:
                    t+=1
                while t>k:
                    if answerKey[l]==s:
                        t-=1
                    l+=1
                ans_T=max(ans_T,r-l+1)
            return ans_T
            
        return max(check('T'),check('F'))
