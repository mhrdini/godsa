package graphs

import (
	"github.com/mhrdini/godsa/datastructures/containers"
)

type Graph interface {
	containers.Container[int]
	Adjacent(v1, v2 int) bool
	Neighbors(v int) []int
	AddVertex()
	RemoveVertex(v int) bool
	AddEdge(src, dst, weight int) (ok bool)
	UpdateEdge(src, dst, weight int) (inserted bool)
	RemoveEdge(src, dst int) (ok bool)
}

type Options struct {
	TotalVertices uint32
	Undirected    bool
}
