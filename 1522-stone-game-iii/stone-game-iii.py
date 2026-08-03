class Solution:
    def stoneGameIII(self, stoneValue: List[int]) -> str:
        n= len(stoneValue)
        dp = [float('-inf')] * (n + 1)
        dp[n] = 0
        
        for i in range(n - 1, -1, -1):
            take_sum = 0
            for k in range(1, 4):
                if i + k > n:
                    break
                take_sum += stoneValue[i + k - 1]
                dp[i] = max(dp[i], take_sum - dp[i + k])
                
        if dp[0] > 0:
            return "Alice"
        elif dp[0] < 0:
            return "Bob"
        else:
            return "Tie"