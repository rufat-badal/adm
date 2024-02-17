package main

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/graph"
)

func main() {
	n := 30
	pruefer := make([]int, n-2)
	for i := 0; i < len(pruefer); i++ {
		pruefer[i] = rand.Intn(n)
	}
	fmt.Printf("Pruefer sequence: %v\n", pruefer)
	g, _ := graph.NewTreeFromPruefer(pruefer)
	fmt.Printf("recovered Pruefer sequence: %v\n", graph.PrueferFromTree(g))
}
