package coordinate

type Interface interface {
	Distance(Interface) int
	GetNeighbors() []Interface
}
