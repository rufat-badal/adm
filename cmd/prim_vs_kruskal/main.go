package main

import (
	"fmt"
	"time"

	"github.com/rufat-badal/adm/graph"
)

func timeMinPrimVsKruskal(edgeProb float64) {
	const nvertices = 1000
	const maxWeight = 100
	const start = 0
	const numExperiments = 10
	g := graph.NewRandomGraph(nvertices, edgeProb, false, maxWeight)
	fmt.Printf("g has %v vertices and %v edges: ", g.NumVertices, g.NumEdges)
	timeStart := time.Now()
	for i := 0; i < numExperiments; i++ {
		graph.MinSpanTreePrim(g, start)
	}
	timeElapsed := time.Since(timeStart)
	fmt.Printf("Prim took %v, ", timeElapsed/numExperiments)
	timeStart = time.Now()
	for i := 0; i < numExperiments; i++ {
		graph.MinSpanTreeKruskal(g, start)
	}
	timeElapsed = time.Since(timeStart)
	fmt.Printf("Kruskal took %v.\n", timeElapsed/numExperiments)
}

func main() {
	for edgeProb := 0.001; edgeProb < 1.0; edgeProb *= 2 {
		timeMinPrimVsKruskal(edgeProb)
	}
}
