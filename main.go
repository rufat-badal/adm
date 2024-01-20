package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	g := graphs.NewRandomGraph(1000, 0.1, false)
	n := g.NumVertices * (g.NumVertices - 1) / 2
	fmt.Printf("Generate graph with %v edges from a total of %v possible edges\n", g.NumEdges, n)
	fmt.Printf("Statistical edge probability: %v\n", float64(g.NumEdges)/float64(n))
}
