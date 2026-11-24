import math
from typing import List


class Solution:

  def maxLength(self, nums: List[int]) -> int:
    n = len(nums)
    ans = 2 if n >= 2 else 1
    l = 0
    prod = 1

    for r in range(n):
      while math.gcd(prod, nums[r]) > 1:
        prod //= nums[l]
        l += 1

      prod *= nums[r]  # Indented inside the for loop
      ans = max(ans, r - l + 1)  # Indented inside the for loop

    return ans
