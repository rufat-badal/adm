package main

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/graph"
)

func main() {
	n := 20
	pruefer := make([]int, n-2)
	for i := 0; i < len(pruefer); i++ {
		pruefer[i] = rand.Intn(n)
	}
	fmt.Println(pruefer)
	g, _ := graph.NewTreeFromPruefer(pruefer)
	fmt.Println(g.NumVertices)
}
