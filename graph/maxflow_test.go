package graph

import (
	"fmt"
	"math/rand"
	"testing"
)

func randCapacity(low int, high int, r *rand.Rand) func() int {
	return func() int {
		return r.Intn(high-low) + low
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
	rcap := randCapacity(1, maxCapacity+1, r)
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

func (g *Graph) addPathEdges(tail int, head int, rcap func() int, dprob float64, r *rand.Rand) int {
	randc := rcap()
	totalc := randc
	fmt.Printf("adding edge from %v to %v with capacity %v\n", tail, head, randc)
	g.AddEdge(tail, head, randc)
	for r.Float64() < dprob {
		fmt.Printf("adding edge from %v to %v\n", tail, head)
		randc = rcap()
		totalc += randc
		g.AddEdge(tail, head, randc)
	}
	return totalc
}

func (g *Graph) addPath(
	nedges int,
	source int,
	sink int,
	rcap func() int,
	dprob float64,
	r *rand.Rand,
) int {
	tail := source
	var head, pathc int
	if nedges == 1 {
		head = sink
		pathc = g.addPathEdges(tail, head, rcap, dprob, r)
		return pathc
	}

	nextv := g.NumVertices
	g.Edges = append(g.Edges, make([][]Edge, nedges-1)...)
	g.Degree = append(g.Degree, make([]int, nedges-1)...)
	g.NumVertices += nedges - 1
	head = nextv
	nextv++
	pathc = g.addPathEdges(tail, head, rcap, dprob, r)

	var curc int
	for i := 0; i < nedges-1; i++ {
		tail = head
		if i < nedges-2 {
			head = nextv
			nextv++
		} else {
			head = sink
		}
		curc = g.addPathEdges(tail, head, rcap, dprob, r)
		if curc < pathc {
			pathc = curc
		}
	}

	return pathc
}

func TestMaxFlow(t *testing.T) {
	r := rand.New(rand.NewSource(RAND_SEED))
	g := newEmptyGraph(2, true)
	fmt.Println(g.addPath(10, 0, 1, randCapacity(10, 20, r), 0.01, r))
}