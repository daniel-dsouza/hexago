package storage

import "github.com/daniel-dsouza/hexagon/coordinate"

// Storage defines the minimum information needed to be useful
type Storage interface {
	Get(coordinate.Interface) (interface{}, bool)
	Set(coordinate.Interface, interface{})
	Neighbors(coordinate.Interface) []coordinate.Interface
}
