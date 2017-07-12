package algorithm

import (
	"testing"

	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

func TestReachable(t *testing.T) {
	m := make(storage.SimpleMap)
	m.Set(co.Axial{Q: 0, R: 0}, 5)
	m.Set(co.Axial{Q: 0, R: 1}, 4)
	m.Set(co.Axial{Q: -1, R: 1}, 3)
	m.Set(co.NewAxial(0, 2), 3)

	result := Reachable(m, co.NewAxial(0, 0), 4)
	answer := map[co.Interface]interface{}{co.NewAxial(0, 0): nil, co.NewAxial(-1, 1): nil, co.NewAxial(0, 1): nil, co.NewAxial(0, 2): nil}

	if len(result) != len(answer) {
		t.Error("expected [{0 0} {-1 1} {0 1} {0 2}], got ", result)
	}

	for _, x := range result {
		if _, ok := answer[x]; !ok {
			t.Error("expected [{0 0} {-1 1} {0 1} {0 2}], got ", result)
		}
	}
}
