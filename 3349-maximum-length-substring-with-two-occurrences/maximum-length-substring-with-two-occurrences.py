class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        freq=dict()
        l,ans=0,0
        for r in range(len(s)):
            if s[r] in freq:
                freq[s[r]]+=1
            else:
                freq[s[r]]=1
            while freq[s[r]]>2:
                if s[l] in freq:
                    freq[s[l]]-=1
                else:
                    freq[s[l]]=0
                l+=1
            ans=max(ans,r-l+1)
        return ans