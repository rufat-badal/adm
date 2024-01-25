package graphs

import (
	"errors"

	queue "github.com/rufat-badal/adm/queues"
)

type NodeState int

const (
	Undiscovered NodeState = iota
	Discovered
	Processed
)

func dfsSort(g Graph, x int, states []NodeState, parent []int, s *queue.Stack[int]) bool {
	states[x] = Discovered
	for y := range g.Edges[x] {
		if states[y] == Undiscovered {
			parent[y] = x
			dfsSort(g, y, states, parent, s)
		} else if states[y] != Processed {
			return true
		}
	}
	return false
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
			cyclic = dfsSort(g, i, states, parent, &s)
		}
		if cyclic {
			return *new([]int), errors.New("cycle detected, cannot sort cyclic graph")
		}
	}
}
