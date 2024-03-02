package graph

import (
	"fmt"
	"math/rand"
	"testing"
)

func randCapacity(maxCapacity int, r *rand.Rand) func() int {
	return func() int {
		return r.Intn(maxCapacity) + 1
	}
}

func (g *Graph) addEdges(
	tail int,
	head int,
	rcap func() int,
	eprop float64,
	dprop float64,
	r *rand.Rand,
) {
	if r.Float64() < eprop {
		g.AddEdge(tail, head, rcap())
		for r.Float64() < dprop {
			g.AddEdge(tail, head, rcap())
		}
	}
}

func newRandomGeneralGraph(
	nvertices int,
	directed bool,
	edgeProbability float64,
	duplicateProbability float64,
	maxCapacity int,
	r *rand.Rand,
) Graph {
	// This function generates much more general graphs with possibly duplicate edges and bi-directional
	// edges for directed graphs
	g := newEmptyGraph(nvertices, directed)
	rcap := randCapacity(maxCapacity, r)
	for tail := 0; tail < nvertices; tail++ {
		for head := 0; head < nvertices; head++ {
			if tail == head {
				continue
			}
			g.addEdges(tail, head, rcap, edgeProbability, duplicateProbability, r)
			if directed {
				g.addEdges(head, tail, rcap, edgeProbability, duplicateProbability, r)
			}
		}
	}
	return g
}

func TestMaxFlow(t *testing.T) {
	r := rand.New(rand.NewSource(RAND_SEED))
	g := newRandomGeneralGraph(100, true, 0.1, 0.01, 100, r)
	rg := g.ResidualGraph()
	fmt.Println(g.NumEdges)
	fmt.Println(rg.NumVertices)
	parent := rg.BFS(0)
	fmt.Println(parent)
}
