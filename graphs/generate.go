package graphs

import (
	"fmt"
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
	fmt.Println(g.Edges)
	g.addDirectedEdge(tail, head)
	g.Degree[tail]++
	g.NumEdges++
	if !g.Directed {
		g.addDirectedEdge(head, tail)
		g.Degree[head]++
	}
	fmt.Println(g.Edges)
	fmt.Println()
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
	fmt.Println(verticesSorted)
	graph := newEmptyGraph(nvertices, true)
	for i := 0; i < len(verticesSorted)-1; i++ {
		// Assure that verticesSorted is the topological sorting of the graph
		fmt.Printf("adding edge (%v, %v)\n", verticesSorted[i], verticesSorted[i+1])
		graph.AddEdge(verticesSorted[i], verticesSorted[i+1])
		// Add (possibly) further edges depending on the edge probability
		for j := i + 1; j < len(verticesSorted); j++ {
			if rand.Float64() < edgeProbability {
				fmt.Printf("adding edge (%v, %v)\n", verticesSorted[i], verticesSorted[j])
				graph.AddEdge(verticesSorted[i], verticesSorted[j])
			}
		}
	}
	return graph, verticesSorted
}
