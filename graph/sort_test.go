package graph

import (
	"math/rand"
	"slices"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	nvertices := 1000
	sortedWant := r.Perm(nvertices)
	g := NewRandomDAG(sortedWant, 0.1)
	sortedGot, e := TopologicalSort(g)
	if e != nil {
		t.Error(e)
	} else if !slices.Equal(sortedWant, sortedGot) {
		t.Errorf("TopologicalSort returned %v, wanted %v", sortedGot, sortedWant)
	}
	g.addDirectedEdge(sortedWant[nvertices-1], sortedWant[0])
	sortedWant = make([]int, 0)
	sortedGot, e = TopologicalSort(g)
	if e == nil {
		t.Errorf("Topological did not return an error")
	} else if !slices.Equal(sortedGot, sortedWant) {
		t.Errorf("TopologicalSort returned %v, wanted %v", sortedGot, sortedWant)
	}
}
