func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    // Map to store a bitmask of reserved seats for each row
    rowMap := make(map[int]int)
    for _, seat := range reservedSeats {
        row, col := seat[0], seat[1]
        // We only track columns 2 through 9
        if col >= 2 && col <= 9 {
            rowMap[row] |= (1 << (col - 1))
        }
    }

    // Every completely empty row can accommodate 2 groups
    count := (n - len(rowMap)) * 2

    // Correct bitmasks based on (col - 1) shifting:
    // Left block:   cols 2,3,4,5 -> shifts 1,2,3,4
    // Middle block: cols 4,5,6,7 -> shifts 3,4,5,6
    // Right block:  cols 6,7,8,9 -> shifts 5,6,7,8
    left := (1 << 1) | (1 << 2) | (1 << 3) | (1 << 4)
    middle := (1 << 3) | (1 << 4) | (1 << 5) | (1 << 6)
    right := (1 << 5) | (1 << 6) | (1 << 7) | (1 << 8)

    for _, mask := range rowMap {
        lOk := (mask & left) == 0
        rOk := (mask & right) == 0
        
        if lOk {
            count++
        }
        if rOk {
            count++
        }
        // Middle block only counts if neither left nor right blocks could be used
        if !lOk && !rOk && (mask & middle) == 0 {
            count++
        }
    }

    return count
}
