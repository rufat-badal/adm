package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graph"
)

func main() {
	const nvertices = 10
	g := graph.NewRandomGraph(nvertices, 0.5, false, 100)
	fmt.Println(g.NumEdges)
	fmt.Println(graph.TravelingSalesman(g))
}
