class MyQueue:

    def __init__(self):
        self.s=[]
        self.len=0
    
    def push(self, x: int) -> None: 
        self.s.append(x)
        self.len+=1

    def pop(self) -> int:
        self.len-=1
        return self.s.pop(0)
        
        

    def peek(self) -> int:
        return self.s[-self.len]
        

    def empty(self) -> bool:
        return self.len==0
        


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()