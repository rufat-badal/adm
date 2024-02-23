package graph

import (
	"math/rand"
	"testing"

	"github.com/rufat-badal/adm/queue"
)

func shortestPathSimple(tree Graph, start int) ShortestPathResult {
	res := ShortestPathResult{Distance: make([]int, tree.NumVertices), Parent: make([]int, tree.NumVertices)}
	for i := range res.Distance {
		res.Distance[i] = MAXINT
	}
	for i := range res.Parent {
		res.Parent[i] = -1
	}
	res.Distance[start] = 0
	discovered := make([]bool, tree.NumVertices)
	q := queue.NewFIFOQueue[int]()
	q.Enqueue(start)

	for i, e := q.Dequeue(); e == nil; i, e = q.Dequeue() {
		discovered[i] = true
		for _, edge := range tree.Edges[i] {
			j := edge.Head
			if discovered[j] {
				continue
			}
			res.Distance[j] = res.Distance[i] + edge.Weight
			res.Parent[j] = i
			q.Enqueue(j)
		}
	}

	return res
}

func sumEdgeWeightOnPath(weight []int, parent []int, start int, end int) int {
	path := treePath(start, end, parent)
	sum := 0
	for _, i := range path {
		sum += weight[i]
	}
	return sum
}
func TestShortestPathDijkstra(t *testing.T) {
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
	randomizeTreeEdgeWeights(&tree, rng)
	resShould := shortestPathSimple(tree, start)
	g := addRandomEdgesToTree(tree, resShould.Parent, nedges, rng, sumEdgeWeightOnPath)
	res, _ := ShortestPathDijkstra(g, start)
	if e != nil {
		t.Fatal("could not run Dijkstra's algorithm on a valid graph without cycles")
	}
	if len(res.Distance) != len(resShould.Distance) {
		t.Errorf("Distance property of the result has incorrect length %v (should be %v)", len(res.Distance), len(resShould.Distance))
	}
	for i := 0; i < min(len(res.Distance), len(resShould.Distance)); i++ {
		if res.Distance[i] != resShould.Distance[i] {
			t.Errorf("Distance to %v is %v (should be %v)", i, res.Distance[i], resShould.Distance[i])
		}
	}
	if len(res.Parent) != len(resShould.Parent) {
		t.Errorf("Parent property of the result has incorrect length %v (should be %v)", len(res.Parent), len(resShould.Parent))
	}
	for i := 0; i < min(len(res.Parent), len(resShould.Parent)); i++ {
		if res.Parent[i] != resShould.Parent[i] {
			t.Errorf("Parent of %v is %v (should be %v)", i, res.Parent[i], resShould.Parent[i])
		}
	}
}
