# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        
        def suck(node):
            node=node.right
            while node.left:
                node=node.left
            return node
        def dele(node,key):
            if not node:
                return None
            if key>node.val:
                node.right=dele(node.right,key)
            elif key<node.val:
                node.left=dele(node.left,key)
            else:
                if not node.left:
                    return node.right
                if not node.right:
                    return node.left
                s=suck(node)
                node.val=s.val
                node.right=dele(node.right,s.val)
            return node
        return dele(root,key)