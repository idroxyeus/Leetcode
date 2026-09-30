func maxDepthAfterSplit(seq string) []int {
    // 1. Changed len(s) to len(seq)
    answer := make([]int, len(seq)) 
    depth := 0

    // 2. Changed range s to range seq
    for i, char := range seq { 
        if char == '(' {
            answer[i] = depth % 2
            depth++
        } else {
            depth--
            answer[i] = depth % 2
        }
    }

    return answer
}
