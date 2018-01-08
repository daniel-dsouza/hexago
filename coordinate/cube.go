package coordinate

// Cube coordinate with constraint x+y+z = 0
type Cube struct {
	X int
	Y int
	Z int
}

var cubeDirections = [6]Cube{Cube{1, 0, -1}, Cube{0, 1, -1}, Cube{-1, 1, 0}, Cube{-1, 0, 1}, Cube{0, -1, 1}, Cube{1, -1, 0}}

// NewCube creates a new Cube coordinate
func NewCube(x, y, z int) Cube {
	return Cube{
		x,
		y,
		z,
	}
}

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

// LinearInterpolation returns the hexatons (inclusive) between two hexagons
func (a Cube) LinearInterpolation(c Interface) []Interface {
	b:= c.(Cube)
	dist := a.Distance(b)
	line := make([]Interface, dist+1)
	line[0] = a

	for i:=1; i <= dist; i++ {
		line[i] = Cube{
			X: round(float64(a.X) + (float64(b.X) - float64(a.X))*float64(i) / float64(dist)),
			Y: round(float64(a.Y) + (float64(b.Y) - float64(a.Y))*float64(i) / float64(dist)),
			Z: round(float64(a.Z) + (float64(b.Z) - float64(a.Z))*float64(i) / float64(dist)),
		}
	}

	return line
}