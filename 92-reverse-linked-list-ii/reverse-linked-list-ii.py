# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        dummy=ListNode(0,head)
        curr=dummy.next
        leftprev=dummy
        for _ in range(left-1):
            leftprev=curr
            curr=curr.next
        prev=None
        revTail=curr
        for _ in range(right-left+1):
            nxt=curr.next
            curr.next=prev
            prev=curr
            curr=nxt
        revTail.next=curr
        leftprev.next=prev
        return dummy.next