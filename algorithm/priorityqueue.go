package algorithm

import "container/heap"
import "github.com/daniel-dsouza/hexago/coordinate"

type Item struct {
	value    coordinate.Interface
	path     []coordinate.Interface
	cost     int
	priority int
	index    int
}

type PriorityQueue []*Item

// Len implements sort.Interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less implements sort.Interface
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap implements sort.Interface
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push implements heap.Interface
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop implements heap.Interface
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update adjusts the priority of something on the heap
func (pq *PriorityQueue) Update(item *Item, cost, priority int, parent []coordinate.Interface) {
	item.priority = priority
	item.cost = cost
	item.path = parent
	heap.Fix(pq, item.index)
}
