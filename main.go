package main

import (
	"fmt"

	"github.com/rufat-badal/adm/graph"
)

func main() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	graph.Sort[int](s, func(i, j int) bool { return i < j })
	fmt.Println(s)
}
