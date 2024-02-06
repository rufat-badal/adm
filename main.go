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
	fmt.Println(s)
	graph.Sort[int](s, func(x, y int) bool { return x < y })
	fmt.Println(s)
}
