import heapq
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.size=k
        self.heap=[]
        for num in nums:
            heapq.heappush(self.heap,num)
            if len(self.heap)>self.size:
                heapq.heappop(self.heap)
    def add(self, val: int) -> int:
        if len(self.heap)<self.size:    
            heapq.heappush(self.heap,val)
        else:
            heapq.heappushpop(self.heap,val)
        return self.heap[0]



# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)