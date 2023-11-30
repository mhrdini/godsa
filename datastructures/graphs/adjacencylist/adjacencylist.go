package adjacencylist

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/graphs"
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

const adjacencyList = "AdjacencyList"

// assumes a maximum of one edge between vertices in an undirected graph
type Graph struct {
	totalEdges uint32                          // size of a graph
	list       []*singlylinkedlist.List[*edge] // adjacency lists
	undirected bool
}

type edge struct {
	weight int
	src    int
	dst    int
}

func New(o graphs.Options) graphs.Graph {
	return &Graph{
		0,
		emptyList(int(o.TotalVertices)),
		o.Undirected,
	}
}

func (g *Graph) Name() string {
	return adjacencyList
}

func (g *Graph) Size() int {
	return len(g.list)
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
	// return fmt.Sprintf("total edges: %v, total vertices: %v, undirected: %v\n%v", g.totalEdges, len(g.list), g.undirected, g.list)
	return fmt.Sprintf("%v", g.list)
}

func (g *Graph) Reset() {
	g.totalEdges = 0
	g.list = emptyList(len(g.list))
}

func (g *Graph) Adjacent(v1, v2 int) bool {
	_, ok := g.hasEdges(v1, v2)
	return ok
}

func (g *Graph) Neighbors(v int) []int {
	vs := []int{}
	for _, e := range g.list[v].Values() {
		vs = append(vs, e.dst)
	}
	return vs
}

func (g *Graph) Transpose() graphs.Graph {
	if !g.undirected {
		list := emptyList(len(g.list))
		for _, v := range g.list {
			for _, e := range v.Values() {
				edge := &edge{
					weight: e.weight,
					src:    e.dst,
					dst:    e.src,
				}
				list[e.dst].Add(edge)
			}
		}
		return &Graph{
			totalEdges: g.totalEdges,
			list:       list,
			undirected: g.undirected,
		}
	}
	return g
}

func (g *Graph) AddVertex() {
	g.list = append(g.list, singlylinkedlist.New[*edge]())
}

func (g *Graph) RemoveVertex(v int) bool {
	if !g.withinRange(v) {
		return false
	}
	for _, list := range g.list {
		edges := []int{}
		for i, edge := range list.Values() {
			if edge.src == v || edge.dst == v {
				edges = append(edges, i)
			}
			if edge.dst > v {
				edge.dst -= 1
			}
		}
		for i := range edges {
			list.Remove(i)
		}
	}

	g.list = append(g.list[:v], g.list[v+1:]...)

	return true
}

func (g *Graph) AddEdge(src, dst, weight int) bool {
	return g.UpdateEdge(src, dst, weight)
}

func (g *Graph) UpdateEdge(src, dst, weight int) bool {
	var directed, undirected bool

	if idxs, ok := g.hasEdges(src, dst); ok {
		e, _ := g.list[src].Get(idxs[0])
		e.weight = weight
		if g.undirected && src != dst {
			e, _ := g.list[dst].Get(idxs[1])
			e.weight = weight
		}
	} else {
		g.totalEdges++
		e := &edge{weight, src, dst}
		directed = g.list[src].Add(e)
		if g.undirected && src != dst {
			e := &edge{weight, dst, src}
			undirected = g.list[dst].Add(e)
		}
		if src == dst {
			undirected = true
		}
	}
	return g.undirected && directed && undirected || directed
}

func (g *Graph) RemoveEdge(src, dst int) bool {

	if idxs, ok := g.hasEdges(src, dst); ok {
		g.list[src].Remove(idxs[0])
		if g.undirected {
			g.list[dst].Remove(idxs[1])
		}
		g.totalEdges--
		return true
	}

	return false
}

func (g *Graph) hasEdges(src, dst int) ([]int, bool) {
	totalVertices := len(g.list)
	idxs := make([]int, 2)
	var directed, undirected bool // boolean checks for whether there are directed or undirected edges between src and dst

	if src < totalVertices && dst < totalVertices && g.list[src] != nil && g.list[dst] != nil {
		edgesFromSrc := g.list[src].Values()
		for i, edge := range edgesFromSrc {
			if edge.src == src && edge.dst == dst {
				idxs[0] = i
				directed = true
				break
			}
		}

		if g.undirected && src != dst {
			edgesFromDst := g.list[dst].Values()
			for i, edge := range edgesFromDst {
				if edge.src == dst && edge.dst == src {
					idxs[1] = i
					undirected = true
					break
				}
			}
		}

		if src == dst {
			idxs[1] = idxs[0]
			undirected = true
		}
	}
	return idxs, g.undirected && directed && undirected || directed
}

func (e *edge) String() string {
	return fmt.Sprintf("(%v, %v, %v)", e.src, e.dst, e.weight)
}

func (g *Graph) withinRange(v int) bool {
	return v < len(g.list)
}

func emptyList(order int) []*singlylinkedlist.List[*edge] {
	list := make([]*singlylinkedlist.List[*edge], order)
	for i := 0; i < order; i++ {
		list[i] = singlylinkedlist.New[*edge]()
	}
	return list
}
