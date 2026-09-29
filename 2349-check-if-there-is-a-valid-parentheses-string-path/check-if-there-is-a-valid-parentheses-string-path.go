package main

func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    
    // Quick pruning: Path length (m + n - 1) must be even for valid parentheses.
    // Also, must start with '(' and end with ')'
    if (m+n-1)%2 != 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
        return false
    }
    
    // Max possible balance k cannot exceed the max path length
    maxK := m + n
    
    // memo[i][j][k] stores: 0 (unvisited), 1 (false), 2 (true)
    memo := make([][][]int, m)
    for i := range memo {
        memo[i] = make([][]int, n)
        for j := range memo[i] {
            memo[i][j] = make([]int, maxK)
        }
    }
    
    var dfs func(i, j, k int) bool
    dfs = func(i, j, k int) bool {
        // Calculate new balance for the current cell
        if grid[i][j] == '(' {
            k++
        } else {
            k--
        }
        
        // If closing tags exceed opening tags or balance exceeds remaining steps, prune
        if k < 0 || k > (m-1-i)+(n-1-j) {
            return false
        }
        
        // Destination reached
        if i == m-1 && j == n-1 {
            return k == 0
        }
        
        // Return cached result if already calculated
        if memo[i][j][k] != 0 {
            return memo[i][j][k] == 2
        }
        
        // Explore moving right or moving down
        res := false
        if i+1 < m && dfs(i+1, j, k) {
            res = true
        } else if j+1 < n && dfs(i, j+1, k) {
            res = true
        }
        
        // Cache result: 2 for true, 1 for false
        if res {
            memo[i][j][k] = 2
        } else {
            memo[i][j][k] = 1
        }
        
        return res
    }
    
    return dfs(0, 0, 0)
}
