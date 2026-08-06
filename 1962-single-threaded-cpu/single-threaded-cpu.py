import heapq
class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        for i in range(len(tasks)):
            tasks[i].append(i)
        tasks.sort(key=lambda x:x[0])
        res,heap=[],[]
        i,time=0,tasks[0][0]
        while i<len(tasks) or heap:
            while i<len(tasks) and tasks[i][0]<=time:
                heapq.heappush(heap,(tasks[i][1],tasks[i][2]))
                i+=1
            if not heap:
                time=tasks[i][0]
            else:
                pt,index=heapq.heappop(heap)
                time+=pt
                res.append(index)
        return res

