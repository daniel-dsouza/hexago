package algorithm

import (
	"testing"

	"reflect"

	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

func TestReachable(t *testing.T) {
	m := storage.NewSimpleMap()
	m.Set(co.Axial{Q: 0, R: 0}, 5)
	m.Set(co.Axial{Q: 0, R: 1}, 4)
	m.Set(co.Axial{Q: -1, R: 1}, 3)
	m.Set(co.NewAxial(0, 2), 3)

	result := Reachable(m, co.NewAxial(0, 0), 4)
	answer := []co.Interface{co.NewAxial(0, 0), co.NewAxial(-1, 1), co.NewAxial(0, 1), co.NewAxial(0, 2)}

	if !reflect.DeepEqual(result, answer) {
		t.Error("expected [{0 0} {-1 1} {0 1} {0 2}], got ", result)
	}
}
