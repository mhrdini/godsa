package dfs

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
)

func Run(g datastructures.Graph) []*graphs.Vertex {
	return recursiveRun(g)
}

func recursiveRun(g datastructures.Graph) []*graphs.Vertex {
	vertices := make([]*graphs.Vertex, g.Size())
	for i := 0; i < g.Size(); i++ {
		vertices[i] = &graphs.Vertex{Color: graphs.White, Value: i, Parent: nil}
	}
	time := 0
	for i := 0; i < g.Size(); i++ {
		if vertices[i].Color == graphs.White {
			Visit(g, vertices, i, &time)
		}
	}
	return vertices
}

func Visit(g datastructures.Graph, vertices []*graphs.Vertex, u int, time *int) {
	*time++
	discovered := vertices[u]
	discovered.Color = graphs.Gray
	for _, v := range g.Neighbors(discovered.Value) {
		neighbor := vertices[v]
		if neighbor.Color == graphs.White {
			neighbor.Parent = discovered
			Visit(g, vertices, v, time)
		}
	}
	*time++
	discovered.Color = graphs.Black
	discovered.Dist = float64(*time)
	fmt.Println(discovered)
}

func Demo() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 8,
		Undirected:    false,
	})
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(7, 6, 1)

	Run(g)
}

func Demo1() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 6,
		Undirected:    false,
	})
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 1, 1)
	g.AddEdge(4, 3, 1)
	g.AddEdge(5, 5, 1)
	Run(g)
}
