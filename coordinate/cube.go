package coordinate

import "math"

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
	b := c.(Cube)
	dist := a.Distance(b)
	line := make([]Interface, dist+1)
	line[0] = a

	for i := 1; i <= dist; i++ {
		line[i] = Cube{
			X: round(float64(a.X) + (float64(b.X)-float64(a.X))*float64(i)/float64(dist)),
			Y: round(float64(a.Y) + (float64(b.Y)-float64(a.Y))*float64(i)/float64(dist)),
			Z: round(float64(a.Z) + (float64(b.Z)-float64(a.Z))*float64(i)/float64(dist)),
		}
	}

	return line
}

// ComputeDistanceHeading calculates the heading (radians) and distance from one hex to another.
// this may not work...
func (a Cube) ComputeDistanceHeading(end Interface) (float64, float64) {
	// set end relative to start
	adjusted := NewCube(end.(Cube).X-a.X, end.(Cube).Y-a.Y, end.(Cube).Z-a.Z)

	smallopp := float64(abs(adjusted.X)) * math.Sin(math.Pi/6)
	smalladj := float64(abs(adjusted.X)) * math.Cos(math.Pi/6)
	bigadj, heading := 0.0, 0.0

	// two cases
	if (adjusted.X > 0 && adjusted.Y > 0) || (adjusted.X < 0 && adjusted.Y > 0) {
		bigadj = float64(abs(adjusted.Y)) + smallopp
		angle := math.Atan(smalladj / bigadj)
		if adjusted.Y < 0 {
			heading = math.Pi - angle
		} else {
			heading = 2*math.Pi - angle
		}
	} else {
		bigadj = float64(abs(adjusted.Y)) - smallopp
		angle := math.Atan(smalladj / bigadj)
		if adjusted.Y < 0 {
			heading = math.Pi + angle
		} else {
			heading = angle
		}
	}

	return heading, math.Sqrt(bigadj*bigadj + smalladj*smalladj)
}
