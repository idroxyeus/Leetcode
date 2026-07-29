"""
# Definition for a QuadTree node.
class Node:
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight
"""

class Solution:
    def construct(self, grid: List[List[int]]) -> 'Node':
        def cons(r,c,n):
            isSame=True
            first=grid[r][c]
            for i in range(r,r+n):
                for j in range(c,c+n):
                    if grid[i][j]!=first:
                        isSame=False
            if isSame:
                return Node(grid[r][c],True)
            n=n//2
            topleft=cons(r,c,n)
            topright=cons(r,c+n,n)
            bottomleft=cons(r+n,c,n)
            bottomright=cons(r+n,c+n,n)
            return Node(0,False,topleft,topright,bottomleft,bottomright)
        return cons(0,0,len(grid))