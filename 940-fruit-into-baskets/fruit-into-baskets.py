class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        seen=dict()
        l,r,ans=0,0,0
        for r in range(len(fruits)):
            if fruits[r] in seen:
                seen[fruits[r]]+=1
            else:
                seen[fruits[r]]=1
            while len(seen)>2:
                if fruits[l] in seen:
                    seen[fruits[l]]-=1
                    if seen[fruits[l]]==0:
                        del seen[fruits[l]]
                l+=1
            ans=max(ans,r-l+1)
        return ans