package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graphs"
)

func main() {
	h := graphs.NewRandomDAG(100, 0.1)
	fmt.Println(h.NumEdges)
}
