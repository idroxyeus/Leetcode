func countServers(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0]) // Fix: ensure n is len(grid[0])
    row := make([]int, m)
    col := make([]int, n)

    // First pass: count the number of servers in each row and column
    for i := range grid {
        for j, x := range grid[i] {
            row[i] += x
            col[j] += x
        }
    }

    // Second pass: count servers that share a row or column with at least one other server
    for i := range grid {
        for j, x := range grid[i] {
            if x == 1 && (row[i] > 1 || col[j] > 1) {
                ans++
            }
        }
    }

    return ans
}
