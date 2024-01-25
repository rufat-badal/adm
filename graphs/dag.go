package graphs

import queue "github.com/rufat-badal/adm/queues"

type NodeState int

const (
	Undiscovered NodeState = iota
	Discovered
	Processed
)

func dfsSort(g Graph, x int, states []NodeState, parent []int, cycleDetected *bool, s *queue.Stack[int]) {
	states[x] = Discovered
	for y := range g.Edges[x] {
		if states[y] == Undiscovered {
			parent[y] = x
			dfsSort(g, y, states, parent, cycleDetected, s)
		} else if (!g.Directed && )
	}
}

func TopologicalSort(g Graph) []int {
	s := queue.NewStack[int]()
	states := make([]NodeState, g.NumVertices)
	parent := make([]int, g.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
}
