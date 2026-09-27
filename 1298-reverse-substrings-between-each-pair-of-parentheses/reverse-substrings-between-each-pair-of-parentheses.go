func reverseParentheses(s string) string {
    n := len(s)
    pair := make([]int, n)
    stack := []int{}

    // Step 1: Map matching pairs of parentheses
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else if s[i] == ')' {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            pair[i] = j
            pair[j] = i
        }
    }

    // Step 2: Traverse and build the result string
    var sb strings.Builder
    dir := 1 // 1 for forward, -1 for backward

    for i := 0; i < n; i += dir {
        if s[i] == '(' || s[i] == ')' {
            i = pair[i]   // Teleport to the matching bracket
            dir = -dir    // Switch traversal direction
        } else {
            sb.WriteByte(s[i])
        }
    }

    return sb.String()
} // <--- Make sure this final closing brace is included!
