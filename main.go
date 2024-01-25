package main

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	r := rand.New(rand.NewSource(42))
	nvertices := 30
	sortedWant := r.Perm(nvertices)
	g := graphs.NewRandomDAG(sortedWant, 0.1)
	g.AddEdge(sortedWant[20], sortedWant[10])
	sortedGot, e := graphs.TopologicalSort(g)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("want: %v, got: %v", sortedWant, sortedGot)
}
