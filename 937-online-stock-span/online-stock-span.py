class StockSpanner:

    def __init__(self):
        self.p=[]
        self.s=[]

    def next(self, price: int) -> int:
        span=1
        while self.p and  self.p[-1][0]<=price:
                span+=self.p.pop()[1]
        self.p.append((price,span))
        return span


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)