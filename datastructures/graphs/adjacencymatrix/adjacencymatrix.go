package adjacencymatrix

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/graphs"
)

const adjacencyMatrix = "AdjacencyMatrix"
const zeroWeight = 0

type Graph struct {
	totalVertices uint32 // order of a graph
	totalEdges    uint32 // size of a graph
	matrix        [][]int
	undirected    bool
}

func New(o graphs.Options) graphs.Graph {
	return &Graph{
		o.TotalVertices,
		0,
		emptyMatrix(o.TotalVertices),
		o.Undirected,
	}
}

func (g *Graph) Name() string {
	return adjacencyMatrix
}
func (g *Graph) Size() int {
	return int(g.totalVertices)
}
func (g *Graph) Empty() bool {
	return g.totalEdges == 0
}
func (g *Graph) Values() []int {
	vs := []int{}
	for i := 0; i < g.Size(); i++ {
		vs = append(vs, i)
	}
	return vs
}
func (g *Graph) String() string {
	return fmt.Sprintf("%v", g.Values())
}
func (g *Graph) Reset() {
	g.totalEdges = 0
	g.matrix = emptyMatrix(g.totalVertices)
}

func (g *Graph) Adjacent(v1, v2 int) bool {
	return g.hasEdge(v1, v2)
}

func (g *Graph) Neighbors(v int) []int {
	vs := []int{}
	for i := 0; i < int(g.totalVertices); i++ {
		if i != v && g.matrix[v][i] != 0 {
			vs = append(vs, i)
		}
	}
	return vs
}

func (g *Graph) AddEdge(src, dst, weight int) bool {
	return g.UpdateEdge(src, dst, weight)
}

func (g *Graph) UpdateEdge(src, dst, weight int) (inserted bool) {
	if !g.hasEdge(src, dst) {
		inserted = true
		g.totalEdges++
	}
	g.matrix[src][dst] = weight
	if g.undirected {
		g.matrix[dst][src] = weight
	}
	return
}

func (g *Graph) RemoveEdge(src, dst int) (ok bool) {
	if g.hasEdge(src, dst) {
		ok = true
	}
	g.matrix[src][dst] = zeroWeight
	if g.undirected {
		g.matrix[dst][src] = zeroWeight
	}
	return
}

func (g *Graph) hasEdge(src, dst int) bool {
	return g.undirected && g.matrix[dst][src] != zeroWeight && g.matrix[src][dst] != zeroWeight || g.matrix[src][dst] != zeroWeight
}

func emptyMatrix(order uint32) [][]int {
	matrix := make([][]int, order)
	rows := make([]int, order*order)
	for i := range matrix {
		matrix[i], rows = rows[:order], rows[order:]
	}
	return matrix
}
