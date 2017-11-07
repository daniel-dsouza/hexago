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

}