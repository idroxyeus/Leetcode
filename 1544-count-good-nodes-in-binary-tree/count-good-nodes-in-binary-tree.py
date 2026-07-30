# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def cunt(node,m):
            if not node:
                return 0
            res=1 if node.val>=m else 0
            res+=cunt(node.left,max(node.val,m))
            res+=cunt(node.right,max(node.val,m))
            return res
        return cunt(root,root.val)
            