package scc

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	"github.com/mhrdini/godsa/algorithms/graphs/dfs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/stacks/linkedliststack"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
	"github.com/mhrdini/godsa/datastructures/utils/sorter"
)

// Strongly Connected Components using Kosaraju's Algorithm

func Run(g datastructures.Graph) [][]int {
	vertices := dfs.Run(g, g.Values()[0])
	sorter.Sort(vertices, func(a, b *graphs.Vertex) int {
		if a.Dist < b.Dist {
			return comparator.Greater
		} else if a.Dist == b.Dist {
			return comparator.Equal
		} else {
			return comparator.Lesser
		}
	})

	transpose := g.Transpose()
	visited := map[int]bool{}
	components := make([][]int, 0)
	s := linkedliststack.New[int]()
	for _, v := range vertices {
		if done, ok := visited[v.Value]; !ok || !done {
			s.Push(v.Value)
			c := make([]int, 0)
			for v, ok := s.Pop(); ok; v, ok = s.Pop() {
				if done, ok := visited[v]; !ok || !done {
					c = append(c, v)
					visited[v] = true
				}
				for _, u := range transpose.Neighbors(v) {
					if done, ok := visited[u]; !ok || !done {
						s.Push(u)
					}
				}
			}
			components = append(components, c)
		}
	}
	return components
}

func Demo() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 8,
		Undirected:    false,
	})
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 6, 1)
	g.AddEdge(3, 2, 1)
	g.AddEdge(3, 7, 1)
	g.AddEdge(4, 0, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(6, 5, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(7, 7, 1)
	fmt.Println(Run(g))
}
