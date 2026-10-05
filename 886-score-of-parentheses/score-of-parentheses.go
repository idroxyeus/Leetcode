func scoreOfParentheses(s string) int {
    score := 0
    depth := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            depth++
        } else {
            depth--
            // If the previous character was '(', we found a core "()" pair
            if s[i-1] == '(' {
                score += 1 << depth
            }
        }
    }

    return score
}