package graphs

import (
	"math/rand"
)

type Graph struct {
	Edges       [][]int
	Degree      []int
	Directed    bool
	NumVertices int
	NumEdges    int
}

func (g *Graph) addDirectedEdge(tail, head int) {
	g.Edges[tail] = append(g.Edges[tail], head)
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

func newEmptyGraph(nvertices int, directed bool) Graph {
	edges := make([][]int, nvertices)
	degree := make([]int, nvertices)
	return Graph{edges, degree, directed, nvertices, 0}
}

func NewRandomGraph(nvertices int, edgeProbability float64, directed bool) Graph {
	graph := newEmptyGraph(nvertices, directed)
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

func NewRandomDAG(nvertices int, edgeProbability float64) (Graph, []int) {
	// We will create a graph whose vertices have this exact topological sorting:
	verticesSorted := rand.Perm(nvertices)
	graph := newEmptyGraph(nvertices, true)
	for i := 0; i < len(verticesSorted)-1; i++ {
		// Assure that verticesSorted is the topological sorting of the graph
		graph.AddEdge(verticesSorted[i], verticesSorted[i+1])
		// Add (possibly) further edges depending on the edge probability
		for j := i + 1; j < len(verticesSorted); j++ {
			if rand.Float64() < edgeProbability {
				graph.AddEdge(i, j)
			}
		}
	}
	return graph, verticesSorted
}
