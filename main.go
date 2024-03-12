package main

import (
	"fmt"

	"github.com/rufat-badal/adm/queue"
)

func main() {
	// const nvertices = 10
	// g := graph.NewRandomGraph(nvertices, 0.5, false, 100)
	// fmt.Println(g.NumEdges)
	// fmt.Println(graph.TravelingSalesman(g))
	items := make([]queue.HeapItem[[]int], 3)
	items[0] = queue.HeapItem[[]int]{Value: []int{1, 2, 3}, Weight: 10}
	items[1] = queue.HeapItem[[]int]{Value: []int{2, 3}, Weight: 4}
	items[2] = queue.HeapItem[[]int]{Value: []int{2, 3, 42}, Weight: 5}
	fmt.Println(items)
	q := queue.NewMinHeap[[]int](items)
	fmt.Println(q)
	q.Insert(queue.HeapItem[[]int]{Value: []int{35, 36}, Weight: 2})
	fmt.Println(q)
	fmt.Println(q.ExtractMin())
	fmt.Println(q)
}
