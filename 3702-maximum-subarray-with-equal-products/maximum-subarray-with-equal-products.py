from typing import List


class Solution:

  def maxLength(self, nums: List[int]) -> int:
    # Bitmask of prime factors {2: 1, 3: 2, 5: 4, 7: 8} for indices 0..10
    MASKS = (0, 0, 1, 2, 1, 4, 3, 8, 1, 2, 5)

    ans = 2
    l = 0
    mask = 0

    for r, x in enumerate(nums):
      m = MASKS[x]
      while mask & m:
        mask ^= MASKS[nums[l]]
        l += 1
      mask |= m
      if r - l + 1 > ans:
        ans = r - l + 1

    return ans
