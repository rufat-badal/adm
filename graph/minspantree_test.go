package graph

import (
	"math/rand"
	"testing"
)

func minSpanTreeSimple(tree Graph, start int) MinSpanTreeResult {
	// Minimum spanning tree when the input is already a tree.
	parent := BFS(tree)
	weight := 0
	for _, edges := range tree.Edges {
		for _, edge := range edges {
			weight += edge.Weight
		}
	}
	weight /= 2
	return MinSpanTreeResult{Weight: weight, Parent: parent}
}

func compareMinSpanTreeResults(res MinSpanTreeResult, resShould MinSpanTreeResult, t *testing.T) {
	if len(res.Parent) != len(resShould.Parent) {
		t.Errorf("Parent property of the result has incorrect length %v (should be %v)", len(res.Parent), len(resShould.Parent))
	}
	for i := 0; i < len(res.Parent) && i < len(resShould.Parent); i++ {
		if res.Parent[i] != resShould.Parent[i] {
			t.Errorf("Parent of %v is %v, but should be %v", i, res.Parent[i], resShould.Parent[i])
		}
	}
}

func TestMinSpanTreePrim(t *testing.T) {
	const nverts = 10000
	const start = 0
	pruefer := make([]int, nverts-2)
	rng := rand.New(rand.NewSource(RAND_SEED))
	for i := 0; i < len(pruefer); i++ {
		pruefer[i] = rng.Intn(nverts)
	}
	tree, e := NewTreeFromPruefer(pruefer)
	if e != nil {
		t.Fatalf("could not create tree from a valid Pruefer sequence %v", pruefer)
	}
	resShould := minSpanTreeSimple(tree, start)
	res, _ := MinSpanTreePrim(tree, start)
	if e != nil {
		t.Fatal("could not run Prim's algorithm on valid graph without cycles")
	}
	compareMinSpanTreeResults(res, resShould, t)
}
