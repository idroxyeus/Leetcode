class Solution:
    def calPoints(self, operations: List[str]) -> int:
        record=[]
        for i in operations:
            if i=='+':
                record.append(record[-1]+record[-2])
            elif i=='D' and len(record)>0:
                record.append(2*record[-1])
            elif i=='C' and len(record)>0:
                record.pop()
            else:
                record.append(int(i))
            
        return sum(record)
