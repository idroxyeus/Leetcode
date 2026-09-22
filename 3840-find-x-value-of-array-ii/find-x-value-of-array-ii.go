package main

type Node struct {
	Prod   int
	Remain []int
}

type SegmentTree struct {
	Tree []Node
	N    int
	K    int
}

func NewSegmentTree(nums []int, k int) *SegmentTree {
	n := len(nums)
	st := &SegmentTree{
		Tree: make([]Node, 4*n),
		N:    n,
		K:    k,
	}
	st.build(nums, 0, 0, n-1)
	return st
}

func (st *SegmentTree) merge(left, right Node) Node {
	res := Node{
		Prod:   (left.Prod * right.Prod) % st.K,
		Remain: make([]int, st.K),
	}
	
	// Copy left prefix counts
	for i := 0; i < st.K; i++ {
		res.Remain[i] += left.Remain[i]
	}
	
	// Shift and add right prefix counts by the product of the left interval
	for j := 0; j < st.K; j++ {
		if right.Remain[j] > 0 {
			nextRem := (left.Prod * j) % st.K
			res.Remain[nextRem] += right.Remain[j]
		}
	}
	
	return res
}

func (st *SegmentTree) build(nums []int, node, start, end int) {
	if start == end {
		val := nums[start] % st.K
		st.Tree[node] = Node{
			Prod:   val,
			Remain: make([]int, st.K),
		}
		st.Tree[node].Remain[val] = 1
		return
	}
	
	mid := start + (end-start)/2
	st.build(nums, 2*node+1, start, mid)
	st.build(nums, 2*node+2, mid+1, end)
	st.Tree[node] = st.merge(st.Tree[2*node+1], st.Tree[2*node+2])
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
	if start == end {
		v := val % st.K
		st.Tree[node].Prod = v
		for i := 0; i < st.K; i++ {
			st.Tree[node].Remain[i] = 0
		}
		st.Tree[node].Remain[v] = 1
		return
	}
	
	mid := start + (end-start)/2
	if idx <= mid {
		st.update(2*node+1, start, mid, idx, val)
	} else {
		st.update(2*node+2, mid+1, end, idx, val)
	}
	st.Tree[node] = st.merge(st.Tree[2*node+1], st.Tree[2*node+2])
}

func (st *SegmentTree) query(node, start, end, l, r int) Node {
	if l <= start && end <= r {
		return st.Tree[node]
	}
	
	mid := start + (end-start)/2
	if r <= mid {
		return st.query(2*node+1, start, mid, l, r)
	}
	if l > mid {
		return st.query(2*node+2, mid+1, end, l, r)
	}
	
	leftNode := st.query(2*node+1, start, mid, l, mid)
	rightNode := st.query(2*node+2, mid+1, end, mid+1, r)
	return st.merge(leftNode, rightNode)
}

func resultArray(nums []int, k int, queries [][]int) []int {
	st := NewSegmentTree(nums, k)
	ans := make([]int, len(queries))
	
	for i, q := range queries {
		idx, val, start, xi := q[0], q[1], q[2], q[3]
		
		// 1. Perform persistent point update
		st.update(0, 0, st.N-1, idx, val)
		
		// 2. Query the range [start, n-1]
		resNode := st.query(0, 0, st.N-1, start, st.N-1)
		
		// 3. Extract count matching xi
		ans[i] = resNode.Remain[xi]
	}
	
	return ans
}
