package graph

import "github.com/rufat-badal/adm/queue"

const maxUInt = ^uint(0)
const MAXINT = int(maxUInt >> 1)

type MinSpanTreeResult struct {
	weight int
	parent []int
}

func MinSpanTreePrim(g Graph) MinSpanTreeResult {
	weight := 0
	parent := make([]int, g.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	nodes := make([]queue.HeapItem[int], g.NumVertices)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Weight = MAXINT
	}
	nodesHeap := queue.NewMinHeap[primNode](nodes)
}
