package graph

import (
	"github.com/rufat-badal/adm/queue"
)

type ResidualEdge struct {
	Head     int
	Flow     int
	Residual int
}

type VertexPair struct {
	Tail int
	Head int
}

type ResidualGraph struct {
	Edges            [][]ResidualEdge
	vertexPairToEdge map[VertexPair]*ResidualEdge
	NumVertices      int
}

func newEmptyResidualGraph(nvertices int) ResidualGraph {
	return ResidualGraph{make([][]ResidualEdge, nvertices), make(map[VertexPair]*ResidualEdge), nvertices}
}

func (g Graph) ResidualGraph() ResidualGraph {
	// edge weights are interpreted as capacities
	rg := newEmptyResidualGraph(g.NumVertices)
	for tail, edges := range g.Edges {
		for _, e := range edges {
			head := e.Head
			capacity := e.Weight
			forwardPair := VertexPair{tail, head}
			re, ok := rg.vertexPairToEdge[forwardPair]
			if ok {
				// an edge between tail and head was already encountered
				re.Residual += capacity
			} else {
				// add forward residual edge
				rg.Edges[tail] = append(rg.Edges[tail], ResidualEdge{Head: head, Flow: 0, Residual: capacity})
				rg.vertexPairToEdge[forwardPair] = &rg.Edges[tail][len(rg.Edges[tail])-1]
				// add backward residual edge
				rg.Edges[head] = append(rg.Edges[head], ResidualEdge{Head: tail, Flow: 0, Residual: 0})
				backwardPair := VertexPair{head, tail}
				rg.vertexPairToEdge[backwardPair] = &rg.Edges[head][len(rg.Edges[head])-1]
			}
		}
	}

	return rg
}

func (rg ResidualGraph) BFS(start int) []int {
	parent := make([]int, rg.NumVertices)
	for i := range parent {
		parent[i] = -1
	}
	discovered := make([]bool, rg.NumVertices)
	q := queue.NewFIFOQueue[int]()
	q.Enqueue(start)
	discovered[start] = true

	for tail, e := q.Dequeue(); e == nil; tail, e = q.Dequeue() {
		for _, edge := range rg.Edges[tail] {
			if edge.Residual == 0 {
				// do not walk edges that have no residual flow
				// this is the only difference to standard BFS
				continue
			}
			head := edge.Head
			if !discovered[head] {
				parent[head] = tail
				q.Enqueue(head)
				discovered[head] = true
			}
		}
	}

	return parent
}

func (rg ResidualGraph) FindEdge(tail, head int) *ResidualEdge {
	return rg.vertexPairToEdge[VertexPair{tail, head}]
}

func (rg ResidualGraph) pathVolume(parent []int, start int, end int) int {
	if parent[end] == -1 {
		// there is flow from start that can reach end
		return 0
	}

	e := rg.FindEdge(parent[end], end)
	// we assume that pathVolume is called with valid parent slices
	// => e is never nil
	if start == parent[end] {
		return e.Residual
	} else {
		return min(rg.pathVolume(parent, start, parent[end]), e.Residual)
	}
}

func (rg *ResidualGraph) augmentPath(parent []int, start int, end int, volume int) {
	if start == end {
		return
	}
	e := rg.FindEdge(parent[end], end)
	e.Flow += volume
	e.Residual -= volume
	e = rg.FindEdge(end, parent[end])
	e.Residual += volume
	rg.augmentPath(parent, start, parent[end], volume)
}

func (g Graph) MaxFlow(source int, sink int) int {
	rg := g.ResidualGraph()
	parent := rg.BFS(source)
	volume := rg.pathVolume(parent, source, sink)
	mf := 0

	for volume > 0 {
		mf += volume
		rg.augmentPath(parent, source, sink, volume)
		parent = rg.BFS(source)
		volume = rg.pathVolume(parent, source, sink)
	}

	return mf
}
