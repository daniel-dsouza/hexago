package algorithm

import (
	"container/heap"
	"fmt"
	"reflect"

	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

// Heuristic defines a template to make custom heuristics
type Heuristic func(m *storage.Storage, start, end co.Interface) int

func distanceHeuristic(m *storage.Storage, start, end co.Interface) int {
	return start.Distance(end)
}

// AStar implements the A* algorithm for hexagonal coordinates
func AStar(m storage.Storage, h Heuristic, start, end co.Interface) {
	fringe := make(PriorityQueue, 1)
	expanded := make(map[co.Interface]interface{})

	fringe[0] = &Item{
		value:    start,
		path:     []co.Interface{start},
		cost:     0,
		priority: 0,
		index:    0,
	}
	heap.Init(&fringe)

	for len(fringe) > 0 {
		state := heap.Pop(&fringe).(*Item)
		fmt.Println(state.value)

		if _, ok := expanded[state.value]; ok {
			continue
		}

		expanded[state.value] = nil

		if reflect.DeepEqual(state.value, end) {
			fmt.Println("hello")
			fmt.Println(state.path)
			return
		}

		for _, n := range m.Neighbors(state.value) {
			newState := &Item{
				value:    n,
				path:     append(state.path, n),
				cost:     state.cost + 1,
				priority: state.cost + 1 + state.value.Distance(end),
				index:    0,
			}
			heap.Push(&fringe, newState)
		}
	}

	fmt.Println("no path found")
}
