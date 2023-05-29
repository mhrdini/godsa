package dfs

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/stacks/linkedliststack"
)


func RecursiveRun(g datastructures.Graph) {
	vertices := make([]*graphs.Vertex, g.Size())
	for i := 0; i < g.Size(); i++ {
		vertices[i] = &graphs.Vertex{Color: graphs.White, Value: i, Parent: nil}
	}
	for i := 0; i < g.Size(); i++ {
		if vertices[i].Color == graphs.White {
			visit(g, vertices, i)
		}
	}
	fmt.Println(vertices)
}

func visit(g datastructures.Graph, vertices []*graphs.Vertex, u int) {
	discovered := vertices[u]
	discovered.Color = graphs.Gray
	fmt.Println(vertices)
	for _, v := range g.Neighbors(discovered.Value) {
		neighbor := vertices[v]
		if neighbor.Color == graphs.White {
			neighbor.Parent = discovered
			visit(g, vertices, v)
		}
	}
	discovered.Color = graphs.Black
}

func Demo() {
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
	RecursiveRun(g)
}
