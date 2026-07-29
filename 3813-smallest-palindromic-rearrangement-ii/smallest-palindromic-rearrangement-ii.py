import collections
import math

class Solution:
    def smallestPalindrome(self, s: str, k: int) -> str:
        n = len(s)
        count = collections.Counter(s)

        mid_char = ""
        half_count = [0] * 26
        for ch in sorted(count.keys()):
            freq = count[ch]
            if freq % 2 == 1:
                mid_char = ch
            half_count[ord(ch) - ord('a')] = freq // 2
            
        half_len = n // 2

        def count_arrangements(freqs, length):
            res = 1
            rem = length
            for f in freqs:
                if f == 0: continue
                c = math.comb(rem, f)
                res *= c
                if res > k:
                    return k + 1
                rem -= f
            return res

        left = []
        for i in range(half_len):
            found = False
            for c in range(26):
                if half_count[c] == 0:
                    continue
                half_count[c] -= 1
                ways = count_arrangements(half_count, half_len - 1 - i)
                
                if k <= ways:
                    left.append(chr(ord('a') + c))
                    found = True
                    break
                else:
                    k -= ways
                    half_count[c] += 1
            
            if not found:
                return ""
                
        left_str = "".join(left)
        return left_str + mid_char + left_str[::-1]
