class Solution:
    def canAliceWin(self, n: int) -> bool:
        if n<10:
            return False
        sub,alice_turn=10,True
        while n>=sub and sub>=0:
            n-=sub
            sub-=1
            alice_turn=not alice_turn
        return not alice_turn
        
        