

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	
	// Lists to hold coordinates of 1s in both images
	var count1 [][2]int
	var count2 [][2]int
	
	// Step 1: Collect coordinates of all 1s
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if img1[r][c] == 1 {
				count1 = append(count1, [2]int{r, c})
			}
			if img2[r][c] == 1 {
				count2 = append(count2, [2]int{r, c})
			}
		}
	}
	
	// Step 2: Calculate shift vectors and track their frequencies
	// Using a key serialization: deltaRow * 100 + deltaCol
	shifts := make(map[int]int)
	maxOverlap := 0
	
	for _, p1 := range count1 {
		for _, p2 := range count2 {
			dr := p2[0] - p1[0]
			dc := p2[1] - p1[1]
			key := dr*100 + dc
			
			shifts[key]++
			if shifts[key] > maxOverlap {
				maxOverlap = shifts[key]
			}
		}
	}
	
	return maxOverlap
}
