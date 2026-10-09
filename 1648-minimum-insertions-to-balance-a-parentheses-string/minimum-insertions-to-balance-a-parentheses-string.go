func minInsertions(s string) int {
    res := 0
    rightNeeded := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            if rightNeeded%2 == 1 {
                res++
                rightNeeded--
            }
            rightNeeded += 2
        } else {
            rightNeeded--
            if rightNeeded < 0 {
                res++
                rightNeeded += 2
            }
        }
    }

    return res + rightNeeded
}