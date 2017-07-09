package algorithm

import (
	"testing"

	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

func TestAStar(t *testing.T) {
	m := storage.NewSimpleMap()
	m.Set(co.Axial{Q: 0, R: 0}, 5)
	m.Set(co.Axial{Q: 0, R: 1}, 4)
	m.Set(co.Axial{Q: -1, R: 1}, 3)
	m.Set(co.NewAxial(0, 2), 3)

	AStar(m, distanceHeuristic, co.NewAxial(0, 0), co.NewAxial(0, 2))
}
