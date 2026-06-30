class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_p=prices[0]
        max_p=0

        for price in prices:
            min_p = min(min_p, price)
            profit = price-min_p
            max_p = max(max_p, profit)

        return max_p