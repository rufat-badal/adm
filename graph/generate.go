package graph

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/queue"
)

type Edge struct {
	Head   int
	Weight int
}

type Graph struct {
	Edges       [][]Edge
	Degree      []int
	Directed    bool
	NumVertices int
	NumEdges    int
}

func (g Graph) Clone() Graph {
	clonedEdges := make([][]Edge, g.NumVertices)
	for i := range clonedEdges {
		clonedEdges[i] = make([]Edge, len(g.Edges[i]))
		copy(clonedEdges[i], g.Edges[i])
	}
	clonedDegree := make([]int, g.NumVertices)
	copy(clonedDegree, g.Degree)
	return Graph{
		Edges:       clonedEdges,
		Degree:      clonedDegree,
		Directed:    g.Directed,
		NumVertices: g.NumVertices,
		NumEdges:    g.NumEdges,
	}
}

func (g *Graph) addDirectedEdge(tail, head, weight int) {
	g.Edges[tail] = append(g.Edges[tail], Edge{head, weight})
	g.Degree[tail]++
}

func (g *Graph) removeDirectedEdge(tail, head int) {
	for edgeId, edge := range g.Edges[tail] {
		if edge.Head == head {
			g.Edges[tail][edgeId] = g.Edges[tail][len(g.Edges[tail])-1]
			g.Edges[tail] = g.Edges[tail][:len(g.Edges[tail])-1]
			g.Degree[tail]--
		}
	}
}

func (g *Graph) AddEdge(tail, head, weight int) {
	// The caller of this function must check by himself that only new edges are added.
	g.addDirectedEdge(tail, head, weight)
	g.NumEdges++
	if !g.Directed {
		g.addDirectedEdge(head, tail, weight)
	}
}

func (g *Graph) RemoveEdge(tail, head int) {
	g.removeDirectedEdge(tail, head)
	g.NumEdges--
	if !g.Directed {
		g.removeDirectedEdge(head, tail)
	}
}

func newEmptyGraph(nvertices int, directed bool) Graph {
	edges := make([][]Edge, nvertices)
	degree := make([]int, nvertices)
	return Graph{edges, degree, directed, nvertices, 0}
}

func randWeight(maxWeight int) int {
	return rand.Intn(maxWeight) + 1
}

func NewRandomGraph(nvertices int, edgeProbability float64, directed bool, maxWeight int) Graph {
	graph := newEmptyGraph(nvertices, directed)
	// This algorithm can be improved if edgeProbability is small.
	for tail := 0; tail < nvertices-1; tail++ {
		for head := tail + 1; head < nvertices; head++ {
			if rand.Float64() < edgeProbability {
				graph.AddEdge(tail, head, randWeight(maxWeight))
			}
			if directed && rand.Float64() < edgeProbability {
				graph.AddEdge(head, tail, randWeight(maxWeight))
			}
		}
	}
	return graph
}

func NewRandomDAG(sorted []int, edgeProbability float64, maxWeight int) Graph {
	// We will create a graph whose vertices have this exact topological sorting:
	graph := newEmptyGraph(len(sorted), true)
	for i := 0; i < len(sorted)-1; i++ {
		// Assure that verticesSorted is the topological sorting of the graph
		graph.AddEdge(sorted[i], sorted[i+1], randWeight(maxWeight))
		// Add (possibly) further edges depending on the edge probability
		for j := i + 1; j < len(sorted); j++ {
			if rand.Float64() < edgeProbability {
				graph.AddEdge(sorted[i], sorted[j], randWeight(maxWeight))
			}
		}
	}
	return graph
}

func NewTreeFromPruefer(pruefer []int) (Graph, error) {
	nverts := len(pruefer) + 2
	for i, v := range pruefer {
		if v < 0 || v >= nverts {
			return *new(Graph), fmt.Errorf("incorrect vertex id %v in Pruefer sequence at index %v (must be in 0 ... %v)", v, i, nverts-1)
		}
	}

	g := newEmptyGraph(nverts, false)
	degree := make([]int, nverts)
	for i := 0; i < len(degree); i++ {
		degree[i] = 1
	}
	for _, i := range pruefer {
		degree[i]++
	}
	for _, i := range pruefer {
		for j, d := range degree {
			if d == 1 {
				g.AddEdge(i, j, 1)
				degree[i]--
				degree[j]--
				break
			}
		}
	}
	u, v := -1, -1
	for i, d := range degree {
		if d == 1 {
			if u == -1 {
				u = i
			} else {
				v = i
				break
			}
		}
	}
	g.AddEdge(u, v, 1)

	return g, nil
}

func PrueferFromTree(g Graph) []int {
	// This function will modifie the input graph g!
	var leafs []queue.HeapItem[int]
	for i, d := range g.Degree {
		if d == 1 {
			leafs = append(leafs, queue.HeapItem[int]{Value: i, Weight: i})
		}
	}
	leafsHeap := queue.NewMinHeap[int](leafs)
	var pruefer []int
	var i, j int
	nverts := g.NumVertices
	for l, e := leafsHeap.ExtractMin(); e == nil; l, e = leafsHeap.ExtractMin() {
		if nverts == 2 {
			break
		}
		i = l.Value
		j = g.Edges[i][0].Head
		pruefer = append(pruefer, j)
		g.RemoveEdge(i, j)
		if g.Degree[j] == 1 {
			leafsHeap.Insert(queue.HeapItem[int]{Value: j, Weight: j})
		}
		nverts--
	}

	return pruefer
}
