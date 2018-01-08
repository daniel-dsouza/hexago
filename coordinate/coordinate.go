package coordinate

import (
	"math"
)


// Interface defines common methods for coordinates
type Interface interface {
	Distance(Interface) int
	GetNeighbors() []Interface
	LinearInterpolation(Interface) []Interface
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}

	return c
}

func round(a float64) int {
	c := math.Ceil(a)

	if (a > 0 && c-a <= 0.5) || (a < 0 && c-a < 0.5) {
		return int(c)
	}

	return int(math.Floor(a))
}
