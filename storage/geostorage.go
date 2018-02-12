package storage

import (
	"math"

	"github.com/daniel-dsouza/hexago/coordinate"
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
	// attempt using https://knowledge.safe.com/articles/725/calculating-accurate-length-in-meters-for-lat-long.html

	// would i fly around the world with this - no. But for a square mile it should be fine.

	latitudeDifference := latitude - g.lat
	longitudeDiffernece := longitude - g.lon

	verticalOffset := 111131.77741377673104 / math.Pow(1+0.0033584313098335197297*math.Cos(2*latitudeDifference), 1.5)
	horizontalOffset := 111506.26354049367285 * math.Cos(latitudeDifference) / math.Pow(1+0.0033584313098335197297*math.Cos(2*latitudeDifference), 0.5)

	// convert to Axial
	return coordinate.NewAxial(int(g.b11*horizontalOffset), int(g.b21*horizontalOffset+g.b22*verticalOffset))

}

func geographicDistance(lat1, lon1, lat2, lon2 float64) float64 {
	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1

	return math.Sqrt((deltaLat*deltaLat)+(deltaLon+deltaLon)) * 1.113195e5
}

func (g *GeoStorage) geographicOffset(mNorth, mEast float64) (float64, float64) {
	earthRadius := 6378137.0

	deltaLat := mNorth / earthRadius
	deltaLon := mEast / (earthRadius * math.Cos(math.Pi*g.lat/180))

	newLat := g.lat + (deltaLat * 180 / math.Pi)
	newLon := g.lon + (deltaLon * 180 / math.Pi)

	return newLat, newLon
}
