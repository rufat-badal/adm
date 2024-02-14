package graph

import (
	"errors"

	"github.com/rufat-badal/adm/queue"
)

type NodeState int

const (
	Undiscovered NodeState = iota
	Discovered
	Processed
)

func dfsSort(g Graph, x int, states []NodeState, parent []int, s *queue.Stack[int], cyclic *bool) {
	states[x] = Discovered
	for _, edge := range g.Edges[x] {
		y := edge.Head
		if *cyclic {
			return
		}
		if states[y] == Undiscovered {
			parent[y] = x
			dfsSort(g, y, states, parent, s, cyclic)
		} else if states[y] == Discovered {
			*cyclic = true
			return
		}
	}
	states[x] = Processed
	s.Push(x)
}

func TopologicalSort(g Graph) ([]int, error) {
	s := queue.NewStack[int]()
	states := make([]NodeState, g.NumVertices)
	parent := make([]int, g.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	cyclic := false
	for i := 0; i < g.NumVertices; i++ {
		if states[i] == Undiscovered {
			dfsSort(g, i, states, parent, &s, &cyclic)
		}
		if cyclic {
			return *new([]int), errors.New("cycle detected, cannot sort cyclic graph")
		}
	}
	sorted := make([]int, g.NumVertices)
	for i := 0; i < g.NumVertices; i++ {
		x, e := s.Pop()
		if e != nil {
			return *new([]int), errors.New("dfs missed a node of the graph (this should be impossible)")
		}
		sorted[i] = x
	}
	return sorted, nil
}
