package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	g, _ := graphs.NewRandomDAG(5, 0.4)
	fmt.Println(g.NumEdges)
	sorted, e := graphs.TopologicalSort(g)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(sorted)
}
