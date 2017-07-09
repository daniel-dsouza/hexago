package algorithm

import (
	"container/heap"
	"fmt"
	"testing"

	co "github.com/daniel-dsouza/hexagon/coordinate"
)

func TestPriorityQueue(t *testing.T) {
	items := map[co.Interface]int{
		co.NewAxial(0, 5): 5,
		// co.NewAxial(1, 2): 4,
		// co.NewAxial(3, 4): 2,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for v, p := range items {
		pq[i] = &Item{
			value:    v,
			cost: 0,
			priority: p,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    co.NewAxial(2, 7),
		priority: 8,
		index:    7,
	}
	heap.Push(&pq, item)

	pq.Update(item, 3, 5, nil)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%v ", item.priority, item.value)
	}
}
