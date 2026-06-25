class Solution:
    def addTwoNumbers(self, l1, l2):

        val1 = 0
        val2 = 0

        place1 = place2 = 1

        d1 = l1
        d2 = l2

        while d1:
            val1 += d1.val * place1
            place1 *= 10
            d1 = d1.next

        while d2:
            val2 += d2.val * place2
            place2 *= 10
            d2 = d2.next

        res = val1 + val2

        if res == 0:
            return ListNode(0)

        dummy = ListNode(0)
        tail = dummy

        while res:
            digit = res % 10

            tail.next = ListNode(digit)
            tail = tail.next

            res //= 10

        return dummy.next