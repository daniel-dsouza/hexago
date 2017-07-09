package algorithm

import (
	co "github.com/daniel-dsouza/hexagon/coordinate"
	"github.com/daniel-dsouza/hexagon/storage"
)

func Reachable(storage storage.Storage, start co.Interface, movement int) []co.Interface {
	visited := map[co.Interface]interface{}{start: nil}
	fringes := make([][]co.Interface, movement+1, movement+1)
	for i := 1; i <= movement; i++ {
		fringes[i] = make([]co.Interface, 0, 6)
	}

	fringes[0] = append(fringes[0], start)

	for i := 1; i <= movement; i++ {
		for _, c := range fringes[i-1] {
			for _, n := range storage.Neighbors(c) {
				if _, ok := visited[n]; !ok {
					visited[n] = nil
					fringes[i] = append(fringes[i], n)
				}
			}
		}
	}

	keys := make([]co.Interface, len(visited))
	i := 0
	for k := range visited {
		keys[i] = k
		i++
	}

	return keys
}
