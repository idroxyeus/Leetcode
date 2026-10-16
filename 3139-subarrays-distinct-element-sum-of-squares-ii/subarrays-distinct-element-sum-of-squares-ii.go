const mod = 1e9 + 7

type SegmentTree struct {
	n        int
	sum      []int64
	square   []int64
	lazy     []int64
}

func NewSegmentTree(n int) *SegmentTree {
	return &SegmentTree{
		n:      n,
		sum:    make([]int64, 4*n),
		square: make([]int64, 4*n),
		lazy:   make([]int64, 4*n),
	}
}

func (st *SegmentTree) pushDown(node, l, r int) {
	if st.lazy[node] != 0 {
		add := st.lazy[node]
		mid := (l + r) / 2
		st.apply(2*node, l, mid, add)
		st.apply(2*node+1, mid+1, r, add)
		st.lazy[node] = 0
	}
}

func (st *SegmentTree) apply(node, l, r int, add int64) {
	length := int64(r - l + 1)
	// (c + add)^2 = c^2 + 2*c*add + add^2 * length
	st.square[node] = (st.square[node] + 2*st.sum[node]*add + add*add*length) % mod
	st.sum[node] = (st.sum[node] + add*length) % mod
	st.lazy[node] = (st.lazy[node] + add) % mod
}

func (st *SegmentTree) updateRange(node, l, r, ql, qr int, add int64) {
	if ql <= l && r <= qr {
		st.apply(node, l, r, add)
		return
	}
	st.pushDown(node, l, r)
	mid := (l + r) / 2
	if ql <= mid {
		st.updateRange(2*node, l, mid, ql, qr, add)
	}
	if qr > mid {
		st.updateRange(2*node+1, mid+1, r, ql, qr, add)
	}
	st.square[node] = (st.square[2*node] + st.square[2*node+1]) % mod
	st.sum[node] = (st.sum[2*node] + st.sum[2*node+1]) % mod
}

func sumCounts(nums []int) int {
	n := len(nums)
	st := NewSegmentTree(n)
	lastSeen := make(map[int]int)
	var ans int64 // Removed unused totalSum variable

	for r := 0; r < n; r++ {
		l := 0
		if prev, exists := lastSeen[nums[r]]; exists {
			l = prev + 1
		}
		st.updateRange(1, 0, n-1, l, r, 1)
		lastSeen[nums[r]] = r

		// Fixed: accessing the root node (index 1) of the segment tree
		ans = (ans + st.square[1]) % mod 
	}

	return int(ans)
}

