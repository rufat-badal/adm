package graphs

import (
	"log"

	queue "github.com/rufat-badal/adm/queues"
)

func bfsFromNode(graph Graph, discovered []bool, parent []int, start int) {
	q := queue.NewFIFOQueue[int]()
	q.Enqueue(start)
	discovered[start] = true

	for !q.IsEmpty() {
		x, e := q.Dequeue()
		if e != nil {
			log.Fatal(e)
		}
		for y := range graph.Edges[x] {
			if !discovered[y] {
				parent[y] = x
				q.Enqueue(y)
				discovered[y] = true
			}
		}
	}
}

func BFS(graph Graph) []int {
	discovered := make([]bool, graph.NumVertices)
	parent := make([]int, graph.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	for node := 0; node < graph.NumVertices; node++ {
		if !discovered[node] {
			bfsFromNode(graph, discovered, parent, node)
		}
	}
	return parent
}
