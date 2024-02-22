package graph

import (
	"math/rand"
	"testing"
)

func minSpanTreeSimple(tree Graph, start int) MinSpanTreeResult {
	// Minimum spanning tree when the input is already a tree.
	parent := BFS(tree, start)
	weight := 0
	for _, edges := range tree.Edges {
		for _, edge := range edges {
			weight += edge.Weight
		}
	}
	weight /= 2
	return MinSpanTreeResult{Weight: weight, Parent: parent}
}

func randomizePrimTreeEdgeWeights(tree *Graph, rng *rand.Rand) {
	const maxWeight = 5
	parent := BFS(*tree, 0)
	weight := make([]int, tree.NumVertices)
	for i := range weight {
		weight[i] = rng.Intn(maxWeight)
	}
	for i := range tree.Edges {
		for e := range tree.Edges[i] {
			j := tree.Edges[i][e].Head
			if i == parent[j] {
				tree.Edges[i][e].Weight = weight[j]
			} else {
				tree.Edges[i][e].Weight = weight[i]
			}
		}
	}
}

func maxEdgeWeightOnPath(weight []int, parent []int, start int, end int) int {
	var fromStart, fromEnd []int
	for i := start; i != -1; i = parent[i] {
		fromStart = append(fromStart, i)
	}
	for i := end; i != -1; i = parent[i] {
		fromEnd = append(fromEnd, i)
	}
	fork := 0
	for fork+1 < min(len(fromStart), len(fromEnd)) {
		if fromStart[len(fromStart)-1-(fork+1)] != fromEnd[len(fromEnd)-1-(fork+1)] {
			break
		}
		fork++
	}
	fromStart = fromStart[:len(fromStart)-1-fork]
	fromEnd = fromEnd[:len(fromEnd)-1-fork]
	maxWeight := -1
	for _, i := range fromStart {
		if weight[i] > maxWeight {
			maxWeight = weight[i]
		}
	}
	for _, i := range fromEnd {
		if weight[i] > maxWeight {
			maxWeight = weight[i]
		}
	}
	return maxWeight
}

func addRandomEdgesToPrimTree(tree Graph, parent []int, nedges int, rng *rand.Rand) Graph {
	weight := make([]int, len(parent))
	for j := range weight {
		i := parent[j]
		if i == -1 {
			continue
		}
		for _, edge := range tree.Edges[i] {
			if edge.Head == j {
				weight[j] = edge.Weight
				break
			}
		}
	}
	g := tree.Clone()
	nedgesAdded := 0
	var i, j int
outer:
	for nedgesAdded < nedges {
		i = rng.Intn(tree.NumVertices)
		j = rng.Intn(tree.NumVertices)
		if i == j {
			continue
		}
		for _, edge := range g.Edges[i] {
			if edge.Head == j {
				continue outer
			}
		}
		g.AddEdge(i, j, maxEdgeWeightOnPath(weight, parent, i, j)+1)
		nedgesAdded++
	}
	return g
}

func testMinSpanTreeGeneric(t *testing.T, minSpanTree func(g Graph, start int) (MinSpanTreeResult, error)) {
	const nverts = 10000
	const start = 0
	const nedges = 10000

	pruefer := make([]int, nverts-2)
	rng := rand.New(rand.NewSource(RAND_SEED))
	for i := 0; i < len(pruefer); i++ {
		pruefer[i] = rng.Intn(nverts)
	}
	tree, e := NewTreeFromPruefer(pruefer)
	if e != nil {
		t.Fatalf("could not create tree from a valid Pruefer sequence %v", pruefer)
	}
	randomizePrimTreeEdgeWeights(&tree, rng)
	resShould := minSpanTreeSimple(tree, start)
	g := addRandomEdgesToPrimTree(tree, resShould.Parent, nedges, rng)
	res, _ := minSpanTree(g, start)
	if e != nil {
		t.Fatal("could not run Prim's algorithm on valid graph without cycles")
	}
	if res.Weight != resShould.Weight {
		t.Errorf("weight of minimal spanning tree is %v (should be %v)", res.Weight, resShould.Weight)
	}
	if len(res.Parent) != len(resShould.Parent) {
		t.Errorf("Parent property of the result has incorrect length %v (should be %v)", len(res.Parent), len(resShould.Parent))
	}
	for i := 0; i < len(res.Parent) && i < len(resShould.Parent); i++ {
		if res.Parent[i] != resShould.Parent[i] {
			t.Errorf("Parent of %v is %v, but should be %v", i, res.Parent[i], resShould.Parent[i])
		}
	}
}

func TestMinSpanTree(t *testing.T) {
	testMinSpanTreeGeneric(t, MinSpanTreePrim)
	testMinSpanTreeGeneric(t, MinSpanTreeKruskal)
}
