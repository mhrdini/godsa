package toposort

import (
	"fmt"

	"github.com/mhrdini/godsa/algorithms/graphs"
	"github.com/mhrdini/godsa/algorithms/graphs/dfs"
	"github.com/mhrdini/godsa/helpers"

	datastructures "github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/graphs/adjacencylist"
	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
	"github.com/mhrdini/godsa/datastructures/utils/sorter"
)

func SortBFS(g datastructures.Graph) []int {
	inDegrees := make([]int, g.Size())
	for i := 0; i < g.Size(); i++ {
		for _, e := range g.Neighbors(i) {
			inDegrees[e]++
		}
	}

	// enqueue vertices with 0 indegrees
	// dequeue a vertex
	// add to toposort
	// update indegrees of all vertices

	toposort := []int{}
	q := linkedlistqueue.New[int]()
	queued := map[int]bool{}
	toBeQueued := findZeroInDegree(inDegrees)
	for _, v := range toBeQueued {
		if done, ok := queued[v]; !ok || !done {
			queued[v] = true
			q.Enqueue(v)
		}
	}
	for v, ok := q.Dequeue(); ok; v, ok = q.Dequeue() {
		toposort = append(toposort, v)
		updateInDegrees(g, inDegrees, v)
		toBeQueued := findZeroInDegree(inDegrees)
		for _, v := range toBeQueued {
			if done, ok := queued[v]; !ok || !done {
				queued[v] = true
				q.Enqueue(v)
			}
		}
	}
	return toposort
}

func findZeroInDegree(inDegrees []int) []int {
	zeros := []int{}
	for v, d := range inDegrees {
		if d == 0 {
			zeros = append(zeros, v)
		}
	}
	return zeros
}

func updateInDegrees(g datastructures.Graph, inDegrees []int, removed int) {
	inDegrees[removed] = -1
	affected := g.Neighbors(removed)
	for _, v := range affected {
		if inDegrees[v] != -1 {
			inDegrees[v]--
		}
	}
}

func SortDFS(g datastructures.Graph) []int {
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
	fmt.Println(SortBFS(g))
}
