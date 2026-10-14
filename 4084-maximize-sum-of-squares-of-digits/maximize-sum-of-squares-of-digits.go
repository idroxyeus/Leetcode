package main

import "strings"

func maxSumOfSquares(num int, sum int) string {
    // 1. Check if the required sum is impossible to achieve with the given number of digits
    if sum > 9 * num {
        return ""
    }

    // 2. Pre-allocate a byte slice of the required length for optimal performance
    res := make([]byte, num)
    
    // 3. Greedily fill with as many '9's as possible
    nines := sum / 9
    for i := 0; i < nines; i++ {
        res[i] = '9'
    }
    
    // 4. If there's a remainder, place it immediately after the '9's
    idx := nines
    if idx < num {
        remainder := sum % 9
        res[idx] = byte('0' + remainder)
        idx++
    }
    
    // 5. Fill the remaining positions with '0's
    for i := idx; i < num; i++ {
        res[i] = '0'
    }
    
    return string(res)
}
