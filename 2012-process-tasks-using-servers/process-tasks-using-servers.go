package main

import (
	"container/heap"
)

type Server struct {
	weight int
	index  int
	freeAt int64
}

// FreeHeap stores available servers, sorted by weight, then index.
type FreeHeap []*Server

func (h FreeHeap) Len() int      { return len(h) }
func (h FreeHeap) Less(i, j int) bool {
	if h[i].weight == h[j].weight {
		return h[i].index < h[j].index
	}
	return h[i].weight < h[j].weight
}
func (h FreeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Use pointer receiver (*FreeHeap) to correctly modify the underlying slice
func (h *FreeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Server))
}
func (h *FreeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// BusyHeap stores active servers, sorted by freeAt time, then weight, then index.
type BusyHeap []*Server

func (h BusyHeap) Len() int      { return len(h) }
func (h BusyHeap) Less(i, j int) bool {
	if h[i].freeAt == h[j].freeAt {
		if h[i].weight == h[j].weight {
			return h[i].index < h[j].index
		}
		return h[i].weight < h[j].weight
	}
	return h[i].freeAt < h[j].freeAt
}
func (h BusyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Use pointer receiver (*BusyHeap) to correctly modify the underlying slice
func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(*Server))
}
func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func assignTasks(servers []int, tasks []int) []int {
	n, m := len(servers), len(tasks)
	ans := make([]int, m)

	freeHeap := &FreeHeap{}
	busyHeap := &BusyHeap{}
	heap.Init(freeHeap)
	heap.Init(busyHeap)

	for i := 0; i < n; i++ {
		heap.Push(freeHeap, &Server{weight: servers[i], index: i, freeAt: 0})
	}

	var currTime int64 = 0

	for j := 0; j < m; j++ {
		if int64(j) > currTime {
			currTime = int64(j)
		}

		// To peek at the minimum value safely in Go, cast the pointer to the underlying slice
		if freeHeap.Len() == 0 && busyHeap.Len() > 0 {
			nextFreeTime := (*busyHeap)[0].freeAt
			if nextFreeTime > currTime {
				currTime = nextFreeTime
			}
		}

		// Move all servers that have finished by currTime back to freeHeap
		for busyHeap.Len() > 0 && (*busyHeap)[0].freeAt <= currTime {
			s := heap.Pop(busyHeap).(*Server)
			heap.Push(freeHeap, s)
		}

		// Assign task j to the best available server
		s := heap.Pop(freeHeap).(*Server)
		ans[j] = s.index
		s.freeAt = currTime + int64(tasks[j])
		heap.Push(busyHeap, s)
	}

	return ans
}
