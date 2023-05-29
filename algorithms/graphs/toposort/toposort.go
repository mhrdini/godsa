package toposort

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	"github.com/mhrdini/godsa/algorithms/graphs/dfs"
	"github.com/mhrdini/godsa/helpers"

	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/utils/sorter"
)

func Sort(g datastructures.Graph) []int {
	vertices := dfs.Run(g)
	sorter.Sort(vertices, func(a, b *graphs.Vertex) int {
		if a.Dist < b.Dist {
			return 1
		} else if a.Dist == b.Dist {
			return 0
		} else {
			return -1
		}
	})
	return helpers.Map(vertices, func(v *graphs.Vertex) int {
		return v.Value
	})
}

func Demo() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 9,
		Undirected:    false,
	})
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 6, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(5, 7, 1)
	g.AddEdge(6, 8, 1)
	g.AddEdge(7, 8, 1)
	fmt.Println(Sort(g))
}
