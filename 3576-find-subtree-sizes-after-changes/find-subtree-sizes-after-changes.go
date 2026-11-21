func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[parent[i]] = append(adj[parent[i]], i)
	}

	ans := make([]int, n)
	lastSeen := make([]int, 26)
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	var dfs func(u int)
	dfs = func(u int) {
		ans[u] = 1
		charIdx := s[u] - 'a'
		prevAncestor := lastSeen[charIdx]

		lastSeen[charIdx] = u

		for _, v := range adj[u] {
			dfs(v)
		}

		lastSeen[charIdx] = prevAncestor

		if u != 0 {
			finalParent := parent[u]
			if prevAncestor != -1 {
				finalParent = prevAncestor
			}
			ans[finalParent] += ans[u]
		}
	}

	dfs(0)
	return ans
}
