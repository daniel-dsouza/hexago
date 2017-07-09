package coordinate

// Cube coordinate with constraint x+y+z = 0
type Cube struct {
	X int
	Y int
	Z int
}

var cubeDirections = [6]Cube{Cube{1, 0, -1}, Cube{0, 1, -1}, Cube{-1, 1, 0}, Cube{-1, 0, 1}, Cube{0, -1, 1}, Cube{1, -1, 0}}

func (a Cube) add(b Cube) Cube {
	return Cube{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// Distance finds the distance between two cube coordinates
func (a Cube) Distance(c Interface) int {
	b := c.(Cube)
	return max(abs(a.X-b.X), abs(a.Y-b.Y), abs(a.Z-b.Z))
}

// GetNeighbors returns the six neighboring cube coordinates
func (a Cube) GetNeighbors() []Interface {
	neighbors := make([]Interface, 6, 6)

	for index, value := range cubeDirections {
		neighbors[index] = a.add(value)
	}

	return neighbors
}
