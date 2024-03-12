package graph

import (
	"errors"
	"fmt"

	"github.com/rufat-badal/adm/queue"
)

type ShortestPathResult struct {
	Distance []int
	Parent   []int
}

func ShortestPathDijkstra(g Graph, start int) (ShortestPathResult, error) {
	if start < 0 || start > g.NumVertices {
		return *new(ShortestPathResult), fmt.Errorf("%v is not an admissible starting vertex (only 0, ..., %v allowed)", start, g.NumVertices-1)
	}

	res := ShortestPathResult{Distance: make([]int, g.NumVertices), Parent: make([]int, g.NumVertices)}
	for i := range res.Distance {
		res.Distance[i] = MAXINT
	}
	for i := range res.Parent {
		res.Parent[i] = -1
	}
	nodes := make([]queue.HeapItem[int], g.NumVertices)
	for i := range nodes {
		nodes[i] = queue.HeapItem[int]{Value: i, Weight: MAXINT}
	}
	nodesHeap := queue.NewMinHeapWithDecreaseWeight[int](nodes)
	nodesHeap.DecreaseWeight(start, 0)

	decreaseWeight := func(i int, newWeight int) {
		nodesHeap.DecreaseWeight(i, newWeight)
		res.Distance[i] = newWeight
	}
	decreaseWeight(start, 0)
	inTree := make([]bool, g.NumVertices)

	for node, e := nodesHeap.ExtractMin(); e == nil; node, e = nodesHeap.ExtractMin() {
		i, dist := node.Value, node.Weight
		if inTree[i] {
			continue
		}
		if dist == MAXINT {
			return res, errors.New("graph was not connected, some nodes are unreachable")
		}
		inTree[i] = true
		for _, edge := range g.Edges[i] {
			j, weight := edge.Head, edge.Weight
			if inTree[j] {
				continue
			}
			oldDistance := res.Distance[j]
			newDistance := res.Distance[i] + weight
			if newDistance < oldDistance {
				res.Parent[j] = i
				decreaseWeight(j, newDistance)
			}
		}
	}

	return res, nil
}
