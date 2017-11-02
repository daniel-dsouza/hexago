package storage

import (
	"github.com/daniel-dsouza/hexagon/coordinate"
	"math"
)

// GeoStorage is a wrapper to add points based on lat/lon
type GeoStorage struct {
	lat                float64
	lon                float64
	a11, a12, a21, a22 float64 // hex -> cartesian offset
	b11, b12, b21, b22 float64 // cartesian offset -> hex
	store              *Interface
}

// New creates a new GeoStorage centered around latitutude, longitude
func New(latitude float64, longitude float64, store *Interface) GeoStorage {
	return GeoStorage{
		lat: latitude,
		lon: longitude,
		a11: math.Sqrt(3.0) / 2.0, a12: 0, a21: 0.5, a22: 1,
		b11: 2.0 * math.Sqrt(3.0) / 3.0, b12: 0.0, b21: -math.Sqrt(3.0) / 3.0, b22: 1.0,
		store: store,
	}
}

// AxialFromGeo gets an Axial coordinate from the difference the lat,lon center and the arg.
func (g *GeoStorage) AxialFromGeo(latitude, longitude float64) coordinate.Axial {
	// TODO: calculate offset

	verticalOffset := 0.0
	horizontalOffset := 0.0

	// convert to Axial
	return coordinate.NewAxial(int(g.b11*horizontalOffset),
		int(g.b21*horizontalOffset+g.b22*verticalOffset))

}
