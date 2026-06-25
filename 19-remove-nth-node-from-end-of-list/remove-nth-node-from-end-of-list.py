# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy=ListNode(0,head)
        slow=fast=dummy
        for i in range(n+1):
            fast=fast.next
        while fast:
            slow=slow.next
            fast=fast.next
        slow.next=slow.next.next
        return dummy.next

        # while curr:
        #     nxt=curr.next
        #     curr.next=prev
        #     prev=curr
        #     curr.next=nxt

        # x=1
        # rev_head=curr
        # while x<n:
        #     curr=curr.next
        