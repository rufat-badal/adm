package graph

import (
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
	g.AddEdge(tail, head, randc)
	for r.Float64() < dprob {
		randc = rcap()
		totalc += randc
		g.AddEdge(tail, head, randc)
	}
	return totalc
}

func (g *Graph) expand(nverts int) {
	g.Edges = append(g.Edges, make([][]Edge, nverts)...)
	g.Degree = append(g.Degree, make([]int, nverts)...)
	g.NumVertices += nverts
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
	g.expand(nedges - 1)
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

func (g *Graph) addExtraVerticesAndEdges(nverts int, eprob float64, dprob float64, r *rand.Rand, rcap func() int) {
	oldNumVerts := g.NumVertices
	g.expand(nverts)
	for tail := 0; tail < g.NumVertices; tail++ {
		for head := 0; head < g.NumVertices; head++ {
			if tail == head || head < oldNumVerts {
				continue
			}
			if r.Float64() < eprob {
				g.AddEdge(tail, head, rcap())
				for r.Float64() < dprob {
					g.AddEdge(tail, head, rcap())
				}
			}
		}
	}
}

func TestMaxFlow(t *testing.T) {
	const pathLen = 100
	const lowCap = 1
	const highCap = 1001
	const dprob = 0.05
	const npaths = 100
	const source = 0
	const sink = 1
	const eprob = 0.1
	const nextraVerts = 1000

	r := rand.New(rand.NewSource(RAND_SEED))
	g := newEmptyGraph(2, true)
	maxflowShould := 0
	rcap := randCapacity(lowCap, highCap, r)
	for i := 0; i < npaths; i++ {
		maxflowShould += g.addPath(pathLen, source, sink, rcap, dprob, r)
	}
	g.addExtraVerticesAndEdges(nextraVerts, eprob, dprob, r, rcap)
	maxflow := g.MaxFlow(source, sink)
	if maxflow != maxflowShould {
		t.Errorf("wrong max flow %v computed (should be %v)", maxflow, maxflowShould)
	}
}
