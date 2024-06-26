package graph

import (
	"log"

	"github.com/rufat-badal/adm/queue"
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
		for _, edge := range graph.Edges[x] {
			y := edge.Head
			if !discovered[y] {
				parent[y] = x
				q.Enqueue(y)
				discovered[y] = true
			}
		}
	}
}

func BFS(graph Graph, start int) []int {
	discovered := make([]bool, graph.NumVertices)
	parent := make([]int, graph.NumVertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	for i := 0; i < graph.NumVertices; i++ {
		node := (start + i) % graph.NumVertices
		if !discovered[node] {
			bfsFromNode(graph, discovered, parent, node)
		}
	}
	return parent
}
