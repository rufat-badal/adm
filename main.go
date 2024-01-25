package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	g, _ := graphs.NewRandomDAG(5, 0.0)
	sorted, e := graphs.TopologicalSort(g)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(sorted)
}
