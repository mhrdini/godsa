package dfs

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
)

func Run(g datastructures.Graph, src int) []*graphs.Vertex {
	fmt.Printf("Running DFS on the following graph:\n%v\n\n", g)

	vertices := make([]*graphs.Vertex, g.Size())
	for i := 0; i < g.Size(); i++ {
		vertices[i] = &graphs.Vertex{Color: graphs.White, Value: i, Parent: nil}
	}

	time := 0
	Visit(g, vertices, src, &time)
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
	// CP3 4.4 DAG in visualgo.net
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(7, 6, 1)

	Run(g, 0)
}

func Demo1() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 8,
		Undirected:    false,
	})
	// CP3 4.9 in visualgo.net
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(3, 2, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(5, 7, 1)
	g.AddEdge(6, 4, 1)
	g.AddEdge(7, 6, 1)
	Run(g, 0)
}

func Demo2() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 5,
		Undirected:    false,
	})
	// CP3 4.17 DAG in visualgo.net
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 0, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 1)
	Run(g, 0)
}
