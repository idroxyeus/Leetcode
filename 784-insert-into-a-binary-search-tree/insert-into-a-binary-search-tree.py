# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        def inst(root,key):
            if not root:
               return TreeNode(key)
            if key<root.val:
                root.left=inst(root.left,key)
            else:
                root.right=inst(root.right,key)
            return root
        return inst(root,val)