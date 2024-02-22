package main

import (
	"fmt"
	"math/rand"

	"github.com/rufat-badal/adm/graph"
)

func randHeight(nelems int) int {
	s := graph.NewUnionFind(nelems)
	for s.NumSets() > 1 {
		i := rand.Intn(nelems)
		j := rand.Intn(nelems)
		if i == j {
			continue
		}
		s.Union(i, j)
	}
	return s.Height(0)
}

func main() {
	fmt.Println("Hello, world!")
	nelems := 1
	nexperiments := 10
	nsteps := 6
	for i := 0; i < nsteps; i++ {
		nelems *= 10
		height := 0
		for j := 0; j < nexperiments; j++ {
			height += randHeight(nelems)
		}
		avarageHeight := float64(height) / float64(nexperiments)
		fmt.Printf("Ran %v experiments with nelems = %v: avarage height = %v\n", nexperiments, nelems, avarageHeight)
	}
}
