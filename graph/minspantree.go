package graph

import (
	"errors"
	"fmt"

	"github.com/rufat-badal/adm/queue"
)

const maxUInt = ^uint(0)
const MAXINT = int(maxUInt >> 1)
const MININT = -MAXINT - 1

type MinSpanTreeResult struct {
	Weight int
	Parent []int
}

func MinSpanTreePrim(g Graph, start int) (MinSpanTreeResult, error) {
	if start < 0 || start >= g.NumVertices {
		return *new(MinSpanTreeResult), fmt.Errorf("%v is not an admissible starting vertex (only 0, ..., %v allowed)", start, g.NumVertices-1)
	}
	res := MinSpanTreeResult{0, make([]int, g.NumVertices)}
	for i := 0; i < len(res.Parent); i++ {
		res.Parent[i] = -1
	}
	// heap item weight is the minimimal distance to tree vertices
	// heap item value is node index
	nodes := make([]queue.HeapItem[int], g.NumVertices)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Value = i
		nodes[i].Weight = MAXINT
	}
	nodesHeap := queue.NewMinHeapWithDecreaseWeight[int](nodes)
	inTree := make([]bool, g.NumVertices)
	nodesHeap.DecreaseWeight(0, 0)

	var v, dist int

	for n, e := nodesHeap.ExtractMin(); e == nil; n, e = nodesHeap.ExtractMin() {
		if inTree[n.Value] {
			continue
		}
		if n.Weight == MAXINT {
			return res, errors.New("graph was not connected, the returned minimum spanning tree does not contain all nodes")
		}
		v, dist = n.Value, n.Weight
		inTree[v] = true
		res.Weight += dist
		for _, edge := range g.Edges[v] {
			if inTree[edge.Head] {
				continue
			}
			oldDist := nodesHeap.DecreaseWeight(edge.Head, edge.Weight)
			if oldDist > edge.Weight {
				res.Parent[edge.Head] = v
			}
		}
	}

	return res, nil
}

type EdgePair struct {
	Tail   int
	Head   int
	Weight int
}

func ToEdgePairSlice(g Graph) []EdgePair {
	var l []EdgePair
	for i, edges := range g.Edges {
		for _, edge := range edges {
			j := edge.Head
			w := edge.Weight
			if i < j {
				l = append(l, EdgePair{i, j, w})
			}
		}
	}
	return l
}

func MinSpanTreeKruskal(g Graph, start int) (MinSpanTreeResult, error) {
	if start < 0 || start >= g.NumVertices {
		return *new(MinSpanTreeResult), fmt.Errorf("%v is not an admissible starting vertex (only 0, ..., %v allowed)", start, g.NumVertices-1)
	}

	edges := ToEdgePairSlice(g)
	Sort[EdgePair](edges, func(x, y EdgePair) bool { return x.Weight < y.Weight })
	uf := NewUnionFind(g.NumVertices)
	weight := 0

	tree := newEmptyGraph(g.NumVertices, g.Directed)
	for _, edge := range edges {
		if !uf.SameComponent(edge.Tail, edge.Head) {
			tree.AddEdge(edge.Tail, edge.Head, edge.Weight)
			weight += edge.Weight
			uf.Union(edge.Tail, edge.Head)
		}
	}

	res := MinSpanTreeResult{Weight: weight, Parent: BFS(tree, start)}

	if uf.NumSets() > 1 {
		return res, errors.New("graph was not connected, the returned minimum spanning tree does not contain all nodes")
	}

	return res, nil
}
