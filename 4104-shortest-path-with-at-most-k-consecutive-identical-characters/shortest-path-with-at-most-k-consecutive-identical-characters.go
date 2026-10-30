package main

import (
	"container/heap"
	"math"
)

// Edge represents a directed edge in the graph
type Edge struct {
	to     int
	weight int
}

// State represents an item in the priority queue
type State struct {
	cost  int
	node  int
	count int // number of consecutive identical characters ending at this node
	index int // required for container/heap implementation
}

// PriorityQueue implements heap.Interface and holds States
type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	// 1. Build the adjacency list
	graph := make([][]Edge, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], Edge{to: v, weight: w})
	}

	// 2. Initialize the distance array
	// dist[u][c] stores the minimum cost to reach node `u` with a trailing identical character run of `c`
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	// 3. Initialize Dijkstra's components
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Node 0 starts with a consecutive run length of 1
	dist[0][1] = 0
	heap.Push(pq, &State{cost: 0, node: 0, count: 1})

	// 4. Run Dijkstra processing loop
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*State)
		u, cost, count := curr.node, curr.cost, curr.count

		// Early exit condition if we reached the target destination
		if u == n-1 {
			return cost
		}

		// Skip outdated states
		if cost > dist[u][count] {
			continue
		}

		// Explore outgoing neighbors
		for _, edge := range graph[u] {
			v, weight := edge.to, edge.weight
			nextCount := 1
			
			if labels[u] == labels[v] {
				nextCount = count + 1
			}

			// If the run length exceeds k, this transition is invalid
			if nextCount > k {
				continue
			}

			// Relaxation step
			if dist[u][count]+weight < dist[v][nextCount] {
				dist[v][nextCount] = dist[u][count] + weight
				heap.Push(pq, &State{
					cost:  dist[v][nextCount],
					node:  v,
					count: nextCount,
				})
			}
		}
	}

	return -1
}
