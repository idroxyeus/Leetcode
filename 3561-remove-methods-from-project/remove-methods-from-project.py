class Solution:
    def remainingMethods(self, n: int, k: int, invocations: List[List[int]]) -> List[int]:
        g = [[] for _ in range(n)]
        for u, v in invocations:
            g[u].append(v)
        
        suspicious = [False] * n
        q = deque([k])
        suspicious[k] = True
        
        while q:
            u = q.popleft()
            for v in g[u]:
                if not suspicious[v]:
                    suspicious[v] = True
                    q.append(v)
                    
        for u in range(n):
            if not suspicious[u]:
                for v in g[u]:
                    if suspicious[v]:
                        return list(range(n))
                        
        return [i for i in range(n) if not suspicious[i]]
