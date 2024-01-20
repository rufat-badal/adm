package graphs

import (
	"math/rand"
)

type EdgeNode struct {
	Head int
	Next *EdgeNode
}

type Graph struct {
	Edges       []*EdgeNode
	Degree      []int
	Directed    bool
	NumVertices int
	NumEdges    int
}

func (g *Graph) addDirectedEdge(tail, head int) {
	e := EdgeNode{head, g.Edges[tail]}
	g.Edges[tail] = &e
}

func (g *Graph) AddEdge(tail, head int) {
	// The caller of this function must check by himself that only new edges are added.
	g.addDirectedEdge(tail, head)
	g.Degree[tail]++
	g.NumEdges++
	if !g.Directed {
		g.addDirectedEdge(head, tail)
		g.Degree[head]++
	}
}

func NewRandomGraph(nvertices int, edgeProbability float64, directed bool) Graph {
	edges := make([]*EdgeNode, nvertices)
	degree := make([]int, nvertices)
	graph := Graph{edges, degree, directed, nvertices, 0}
	// This algorithm can be improved if edgeProbability is small.
	for tail := 0; tail < nvertices-1; tail++ {
		for head := tail + 1; head < nvertices; head++ {
			if rand.Float64() < edgeProbability {
				graph.AddEdge(tail, head)
			}
			if directed && rand.Float64() < edgeProbability {
				graph.AddEdge(head, tail)
			}
		}
	}
	return graph
}
