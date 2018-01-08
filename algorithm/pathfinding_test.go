package algorithm

import (
	"testing"

	"reflect"

	co "github.com/daniel-dsouza/hexago/coordinate"
	"github.com/daniel-dsouza/hexago/storage"
)

func TestAStar(t *testing.T) {
	m := storage.NewSimpleMap()
	m.Set(co.Axial{Q: 0, R: 0}, 5)
	m.Set(co.Axial{Q: 0, R: 1}, 4)
	m.Set(co.Axial{Q: -1, R: 1}, 3)
	m.Set(co.NewAxial(0, 2), 3)

	answer := []co.Interface{co.Axial{Q: 0, R: 0}, co.Axial{Q: 0, R: 1}, co.Axial{Q: 0, R: 2}}
	result, ok := AStar(m, distanceHeuristic, co.NewAxial(0, 0), co.NewAxial(0, 2))
	if !ok || !reflect.DeepEqual(answer, result) {
		t.Error("Expected [{0 0} {0 1} {0 2}], got ", result)
	}
}
