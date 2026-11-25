from collections import Counter
from typing import List


class Solution:

  def smallestUniqueSubarray(self, nums: List[int]) -> int:
    n = len(nums)
    MOD = (1 << 61) - 1
    BASE = 1315423911

    powers = [1] * (n + 1)
    prefix = [0] * (n + 1)
    for i in range(n):
      powers[i + 1] = (powers[i] * BASE) % MOD
      prefix[i + 1] = (prefix[i] * BASE + nums[i]) % MOD

    def has_unique(length: int) -> bool:
      counts = Counter()
      p = powers[length]
      for i in range(n - length + 1):
        h = (prefix[i + length] - prefix[i] * p) % MOD
        counts[h] += 1
      return any(c == 1 for c in counts.values())

    low, high = 1, n
    ans = n
    while low <= high:
      mid = (low + high) // 2
      if has_unique(mid):
        ans = mid
        high = mid - 1
      else:
        low = mid + 1

    return ans
