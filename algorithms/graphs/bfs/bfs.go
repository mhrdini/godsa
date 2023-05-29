package bfs

import (
	"fmt"
	"math"

	"github.com/mhrdini/godsa/algorithms/graphs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
)

func Run(g datastructures.Graph, s int) []*graphs.Vertex {
	vertices := make([]*graphs.Vertex, g.Size())
	for i := 0; i < g.Size(); i++ {
		vertices[i] = &graphs.Vertex{Color: graphs.White, Value: i, Dist: math.Inf(1), Parent: nil}
	}
	vertices[s].Color = graphs.Gray
	vertices[s].Dist = 0
	vertices[s].Parent = nil
	q := linkedlistqueue.New[*graphs.Vertex]()
	q.Enqueue(vertices[s])
	for !q.Empty() {
		u, _ := q.Dequeue()
		visited := vertices[u.Value]
		for _, v := range g.Neighbors(u.Value) {
			neighbor := vertices[v]
			if neighbor.Color == graphs.White {
				neighbor.Color = graphs.Gray
				neighbor.Dist = visited.Dist + 1
				neighbor.Parent = visited
				q.Enqueue(neighbor)
			}
		}
		visited.Color = graphs.Black
	}
	return vertices
}

func Demo() {
	g := adjacencylist.New(datastructures.Options{
		TotalVertices: 9,
		Undirected:    true,
	})
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 5, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 7, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 7, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(5, 8, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 1)
	fmt.Println(Run(g, 1))
}
