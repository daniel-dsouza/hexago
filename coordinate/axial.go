package coordinate


// Axial coordinate with implied z, under x+y+z=0
type Axial struct {
	Q int
	R int
}

var axialDirections = [6]Axial{Axial{1, 0}, Axial{1, -1}, Axial{0, -1}, Axial{-1, 0}, Axial{-1, 1}, Axial{0, 1}}

// GetAxialDirections returns the six axial directions starting from 0 degrees counter-clockwise
// func GetAxialDirections() [6]Axial {
// 	return axialDirections
// }

// NewAxial creates a new Axial
func NewAxial(q, r int) Axial {
	return Axial{
		q,
		r,
	}
}

func (a Axial) add(b Axial) Axial {
	return Axial{
		a.Q + b.Q,
		a.R + b.R,
	}
}

//Distance finds the cube distance between two coordinates
func (a Axial) Distance(c Interface) int {
	b := c.(Axial)
	return max(abs(a.Q-b.Q), abs(a.R-b.R), abs(b.Q+b.R-a.Q-a.R))
}

// GetNeighbors returns the six neigboring coordinates
func (a Axial) GetNeighbors() []Interface {
	neighbors := make([]Interface, 6, 6)

	for index, value := range axialDirections {
		neighbors[index] = a.add(value)
	}

	return neighbors
}

// LinearInterpolation returns the hexagons (inclusive) between two hexagons
func (a Axial) LinearInterpolation(c Interface) []Interface {
	b := c.(Axial)
	dist := a.Distance(b)
	line := make([]Interface, dist+1)
	line[0] = a

	for i := 1; i <= dist; i++ {
		line[i] = Axial{
			Q: round(float64(a.Q) + (float64(b.Q)-float64(a.Q))*float64(i)/float64(dist)),
			R: round(float64(a.R) + (float64(b.R)-float64(a.R))*float64(i)/float64(dist)),
		}
	}

	return line
}