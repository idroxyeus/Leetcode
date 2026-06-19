class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        s=[]
        for i in asteroids:
            alive=True
            while alive and s and s[-1]>0 and i<0:
                if s[-1]<-i:
                    s.pop()
                elif s[-1]==-i:
                    s.pop()
                    alive=False
                else:
                    alive=False
            if alive:
                s.append(i)
        return s
                

