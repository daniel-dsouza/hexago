package storage

import "github.com/daniel-dsouza/hexagon/coordinate"

// Interface defines the minimum information needed to be useful
type Interface interface {
	Get(coordinate.Interface) (interface{}, bool)
	Set(coordinate.Interface, interface{})
	LinearInterpolation(coordinate.Interface) []coordinate.Interface
	Neighbors(coordinate.Interface) []coordinate.Interface
}
