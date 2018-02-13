package algorithm

import (
	"github.com/daniel-dsouza/hexago/storage"
	"fmt"
	"github.com/daniel-dsouza/hexago/coordinate"
)

//computeOffset computes a new lat and lon given a heading (radians) and distance (m)
func computeOffset(lat, lon int64, radians, distance float64) (int64, int64) {

}

//coputeOffset computes a new lat and lon given an offset N, E in meters
func computeOffset(lat, lon int64, north, east, float64) {

}

//GeoCell maps a latitude and longitude to a hexagonal cell
struct GeoCell {
	coordinate.Interface,
	latitude int64,
	longitude int64,	
}

//GetGeoNeigbors is based on Axial.GetNeighbors, but also computes the lat/lon
func (GeoCell geo) GetGeoNeighbors() []Interface {
	geoNeighbors := make([]GeoCell, 6, 6)
	neighbors := geo.GetNeighbors()

	for index, value := range geo.GetNeighbors() {
		geoNeighbors[index] = GeoCell {
			value,
			// latitude TODO: use geo.latitude and offset it.
			// longitude TODO: use geo.longitude and offset it.
		}
	}
}

//PopulateMap returns a simplemap with all withing north <-> south, and west <-> east, starting from centerLat, centerLon
func PopulateMap(centerLat, centerLon, northLat, southLat, westLon, eastLon) storage.SimpleMap {
	center := GeoCell {
		coor: coordinate.NewAxial(0, 0)},
		latitude: centerLat,
		longitude: centerLon,
	}
	store := storage.NewSimpleMap()
	store.Set(center, false)

	// time for some RECURSION
	
	
}

