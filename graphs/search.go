package graphs

type NodeState int

const (
	Undiscovered NodeState = iota
	Discovered
	Processed
)

func bfsFromNode(graph Graph, state *[]NodeState, parent *[]int, node int) {
}

func BFS(graph Graph) []int {
	state := make([]NodeState, graph.NumVertices)
	parent := make([]int, graph.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	for node := 0; node < graph.NumVertices; node++ {
		if state[node] != Undiscovered {
			continue
		}
		bfsFromNode(graph, &state, &parent, node)

	}
	return parent
}
