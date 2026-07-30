# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def val(node,left,right):
            if not node:
                return True
            if not (left<node.val<right):
                return False
            l=val(node.left,left,node.val)
            r=val(node.right,node.val,right)
            return l and r
        return val(root,float(-inf),float(inf))
            