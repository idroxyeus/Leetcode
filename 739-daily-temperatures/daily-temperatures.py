class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        s=[]
        n=len(temperatures)
        answer=[0]*n
        for i in range(n):
            while s and temperatures[i]>temperatures[s[-1]]:
                prev=s.pop()
                answer[prev]=i-prev
            s.append(i)
        return answer