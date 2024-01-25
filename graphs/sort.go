package graphs

import (
	"errors"
	"fmt"

	queue "github.com/rufat-badal/adm/queues"
)

type NodeState int

const (
	Undiscovered NodeState = iota
	Discovered
	Processed
)

func dfsSort(g Graph, x int, states []NodeState, parent []int, s *queue.Stack[int], cyclic *bool) {
	states[x] = Discovered
	for y := range g.Edges[x] {
		if *cyclic {
			return
		}
		if states[y] == Undiscovered {
			parent[y] = x
			dfsSort(g, y, states, parent, s, cyclic)
		} else if states[y] == Discovered {
			fmt.Printf("cycle closed from %v to %v\n", x, y)
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
			return *new([]int), errors.New("dfs missed a node of the graph, this should not happen!")
		}
		sorted[i] = x
	}
	return sorted, nil
}
