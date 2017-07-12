package storage

import (
	"github.com/daniel-dsouza/hexagon/coordinate"
)

// SimpleMap is a hashmap representation of a hexagonal grid
type SimpleMap map[coordinate.Interface]interface{}

// NewSimpleMap makes a new simple map
func NewSimpleMap() SimpleMap {
	return make(SimpleMap)
}

// Get returns a the key associated with a coordinate, or the
func (s SimpleMap) Get(key coordinate.Interface) (interface{}, bool) {
	if _, ok := s[key]; ok {
		return nil, false
	}

	return s[key], true
}

// Set associates a value with a coordinate
func (s SimpleMap) Set(key coordinate.Interface, value interface{}) {
	s[key] = value
}

// Neighbors returns the valid neighbors in a map
func (s SimpleMap) Neighbors(key coordinate.Interface) []coordinate.Interface {
	neighbors := make([]coordinate.Interface, 0, 6)

	for _, neighbor := range key.GetNeighbors() {
		if _, ok := s[neighbor]; ok {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
