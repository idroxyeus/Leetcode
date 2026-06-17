class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        pairs = {')': '(', '}': '{', ']': '['}
        for char in s:
            if char in pairs:
                if not stack or stack.pop() != pairs[char]:
                    return False
            else:
                stack.append(char)
        return not stack
# Complexity
# Time: O(n)
# Every bracket is pushed/popped at most once.
# Space: O(n)
# Worst case: "(((((((("

# You cannot do better than O(n) time because every character must be examined at least once
