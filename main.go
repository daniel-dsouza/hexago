package main

import (
	"fmt"

	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

func main() {
	a := co.NewAxial(0, 0)
	b := co.Axial{Q: 0, R: 2}

	fmt.Println(a.Distance(b))
	fmt.Println(b.Distance(a))
	fmt.Println(a.GetNeighbors())
	fmt.Println(b.GetNeighbors())
	fmt.Println(a.GetNeighbors())

	m := storage.NewSimpleMap()
	m.Set(co.Axial{Q: 0, R: 0}, 5)
	m.Set(co.Axial{Q: 0, R: 1}, 4)
	m.Set(co.Axial{Q: -1, R: 1}, 3)
	m.Set(co.NewAxial(0,2), 3)
	fmt.Println(m.Neighbors(a))

	

}
