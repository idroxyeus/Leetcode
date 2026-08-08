class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], minutes: int) -> int:
        unsat,sat=0,0
        for i in range(len(customers)):
            if grumpy[i]==0:
                sat+=customers[i]
        for R in range(minutes):
            if grumpy[R]==1:
                unsat+=customers[R]
        max_un=unsat
        for R in range(minutes, len(customers)):
            if grumpy[R] == 1:
                unsat += customers[R]
            if grumpy[R - minutes] == 1:
                unsat -= customers[R - minutes]
            max_un = max(max_un, unsat)
        return sat+max_un