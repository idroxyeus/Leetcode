class TimeMap:

    def __init__(self):
        self.d={}

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.d:
            self.d[key]=[]
        self.d[key].append([value,timestamp])

    def get(self, key: str, timestamp: int) -> str:
        ans=""
        all_v=self.d.get(key,[])
        l,r=0,len(all_v)-1
        while l<=r:
            m=(l+r)//2
            if all_v[m][1]<=timestamp:
                ans=all_v[m][0]
                l=m+1
            else:
                r=m-1
        return ans
        


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)