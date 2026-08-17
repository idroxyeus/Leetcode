func stoneGameV(stoneValue []int) int {
    n := len(stoneValue)
    prefix := make([]int, n+1)
    for i, v := range stoneValue {
        prefix[i+1] = prefix[i] + v
    }

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dp func(int, int) int
    dp = func(l, r int) int {
        if l == r {
            return 0
        }
        if memo[l][r] != -1 {
            return memo[l][r]
        }

        ans := 0
        for k := l; k < r; k++ {
            s1 := prefix[k+1] - prefix[l]
            s2 := prefix[r+1] - prefix[k+1]

            if s1 < s2 {
                ans = max(ans, s1+dp(l, k))
            } else if s1 > s2 {
                ans = max(ans, s2+dp(k+1, r))
            } else {
                ans = max(ans, s1+max(dp(l, k), dp(k+1, r)))
            }
        }
        memo[l][r] = ans
        return ans
    }

    return dp(0, n-1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
