import heapq
class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        heap=[]
        for weights in stones:
            heapq.heappush(heap,-weights)
        
        
        while len(heap)>1:
            s1,s2=-heapq.heappop(heap),-heapq.heappop(heap)
            if s1==s2:
                continue
            else:
                heapq.heappush(heap,-(s1-s2))
        if not heap:
            return 0
        return -heap[0]

        