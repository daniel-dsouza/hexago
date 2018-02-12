package coordinate

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := NewAxial(45, -63)
	b := NewAxial(-34, 90)
	c := a.add(b)
	d := b.add(a)

	if c != d {
		t.Error("Commutative property not upheld")
	}

	if c != NewAxial(11, 27) || d != NewAxial(11, 27) {
		t.Error("addition went wrong")
	}
}

func TestDistance(t *testing.T) {
	a := NewAxial(32, 78)
	b := NewAxial(12, 34)
	c := NewAxial(-12, 48)

	if a.Distance(b) != 64 {
		t.Error("Distance equation is broken")
	}

	if a.Distance(c) != 74 {
		t.Error("Distance equation is broken for negatives")
	}

	if a.Distance(a) != 0 {
		t.Error("Distance does not hold at 0")
	}

}

func TestNeighbors(t *testing.T) {
	a := NewAxial(1, 1)

	var axialDirections = [6]Axial{Axial{2, 1}, Axial{2, 0}, Axial{1, 0}, Axial{0, 1}, Axial{0, 2}, Axial{1, 2}}

	// result := a.GetNeighbors()

	for i, v := range a.GetNeighbors() {
		if v != axialDirections[i] {
			t.Error("Neighbors do not work")
		}
	}
}

func TestLinearInterpolation(t *testing.T) {
	a := NewAxial(0, 0)
	b := NewAxial(0, 2)

	line := []Axial{Axial{0, 0}, Axial{0, 1}, Axial{0, 2}}

	for i, v := range a.LinearInterpolation(b) {
		if v != line[i] {
			t.Error("Linear interpolation is broken")
		}
	}

}
