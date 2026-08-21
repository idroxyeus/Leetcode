class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res,brac=[],[]
        def  sti(clos,ope):
            if len(brac)==2*n:
                res.append("".join(brac))
                return
            if ope<n:
                brac.append('(')
                sti(clos,ope+1)
                brac.pop()
            if clos<ope:
                brac.append(')')
                sti(clos+1,ope)
                brac.pop()
        sti(0,0)
        return res

                