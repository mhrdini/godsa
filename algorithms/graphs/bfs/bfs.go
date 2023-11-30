package bfs

import (
	"fmt"
	"math"

	"github.com/mhrdini/godsa/algorithms/graphs"
	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
)

func Run(g datastructures.Graph, src int) []*graphs.Vertex {
	vertices := make([]*graphs.Vertex, g.Size())
	for i := 0; i < g.Size(); i++ {
		vertices[i] = &graphs.Vertex{Color: graphs.White, Value: i, Dist: math.Inf(1), Parent: nil}
	}
	vertices[src].Color = graphs.Gray
	vertices[src].Dist = 0
	vertices[src].Parent = nil
	q := linkedlistqueue.New[*graphs.Vertex]()
	q.Enqueue(vertices[src])
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
		fmt.Println(visited)
	}
	return vertices
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

	Run(g, 0)
}
