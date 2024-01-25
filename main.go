package main

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	r := rand.New(rand.NewSource(42))
	nvertices := 10
	sortedWant := r.Perm(nvertices)
	g := graphs.NewRandomDAG(sortedWant, 0.1)
	sortedGot, e := graphs.TopologicalSort(g)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("want: %v, got: %v", sortedWant, sortedGot)
}
