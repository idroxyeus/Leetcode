func stoneGameII(piles []int) int {
    n := len(piles)
    // Suffix sums array to get sum of piles from i to end in O(1)
    sums := make([]int, n)
    sums[n-1] = piles[n-1]
    for i := n - 2; i >= 0; i-- {
        sums[i] = sums[i+1] + piles[i]
    }

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }

    var dp func(int, int) int
    dp = func(i, m int) int {
        if i + 2*m >= n {
            return sums[i]
        }
        if memo[i][m] != 0 {
            return memo[i][m]
        }

        minStones := math.MaxInt32
        for x := 1; x <= 2*m; x++ {
            res := dp(i+x, max(m, x))
            if res < minStones {
                minStones = res
            }
        }

        memo[i][m] = sums[i] - minStones
        return memo[i][m]
    }

    return dp(0, 1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}