func maxDepth(s string) int {
    maxAns := 0
    curr := 0
    
    for _, ch := range s {
        if ch == '(' {
            curr++
            if curr > maxAns {
                maxAns = curr
            }
        } else if ch == ')' {
            curr--
        }
    }
    
    return maxAns
}
