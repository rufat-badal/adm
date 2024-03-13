package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graph"
)

func main() {
	const nvertices = 15
	g := graph.NewRandomGraph(nvertices, 0.5, false, 100)
	fmt.Println(graph.TravelingSalesman(g))
}
