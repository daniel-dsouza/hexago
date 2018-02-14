package coordinate

import (
	"math"
)

//Geo maps a latitude and longitude to a hexagonal cell
type Geo struct {
	Interface
	Latitude  float64
	Longitude float64
}

var earthRadius = 6378137.0 // meters

// computeOffset computes a new lat and lon given a heading (radians) and distance (m)
// based off http://mathforum.org/library/drmath/view/51816.html
func computeRadialOffset(lat, lon, radians, distance float64) (float64, float64) {
	arcLength := distance / earthRadius

	newLat := math.Asin(math.Sin(lat)*math.Cos(arcLength) + math.Cos(lat)*math.Sin(arcLength)*math.Cos(radians))
	deltaLon := math.Atan2(math.Sin(radians)*math.Sin(arcLength)*math.Cos(lat), math.Cos(arcLength)-math.Sin(lat)*math.Sin(newLat))
	newLon := math.Mod(lon-deltaLon+math.Pi, 2*math.Pi) - math.Pi

	return newLat, newLon
}

// computeOffset computes a new lat and lon given an offset N, E in meters
func computeOffset(lat, lon, north, east float64) (float64, float64) {
	deltaLat := north / earthRadius
	deltaLon := east / (earthRadius * math.Cos(math.Pi*lat/180))

	newLat := lat + (deltaLat * 180 / math.Pi)
	newLon := lon + (deltaLon * 180 / math.Pi)

	return newLat, newLon
}

// GetGeoNeighbors is based on Axial.GetNeighbors, but also computes the lat/lon
func (geo Geo) GetGeoNeighbors() []Interface {
	geoNeighbors := make([]Interface, 6, 6)
	angles := []float64{math.Pi / 6, math.Pi / 2, 5 * math.Pi / 6, 7 * math.Pi / 6, 3 * math.Pi / 2, 11 * math.Pi / 6} // FIX ME

	for index, value := range geo.GetNeighbors() {
		newLat, newLon := computeRadialOffset(geo.Latitude, geo.Longitude, angles[index], 2.0)
		geoNeighbors[index] = Geo{
			value,
			newLat, // latitude TODO: use geo.latitude and offset it.
			newLon, // longitude TODO: use geo.longitude and offset it.
		}
	}

	return geoNeighbors
}

// NewGeo creates a new Geo relative to the center hexagon(with radius) of a hexagonal grid
func NewGeo(center Geo, radius float64, coordinate Interface) Geo {
	heading, distance := center.ComputeDistanceHeading(coordinate)
	offsetLat, offsetLon := computeRadialOffset(center.Latitude, center.Longitude, heading, distance)

	return Geo{
		coordinate,
		offsetLat,
		offsetLon,
	}
}
