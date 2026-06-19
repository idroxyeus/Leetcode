class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        s=[]
        for i in tokens:
            if i=='+' or i=='-' or i=='*' or i=='/':
                b=s.pop()
                a=s.pop()
                if i=='+':
                    s.append(a+b)
                elif i=='-':
                    s.append(a-b)
                elif i=='*':
                    s.append(a*b)
                elif i=='/':
                    s.append(int(a/b))
            else:
                s.append(int(i))
        return s.pop()
        