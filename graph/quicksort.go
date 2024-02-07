package graph

import (
	"fmt"
	"math/rand"
)

func Sort[T interface{}](s []T, Less func(x, y T) bool) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	fmt.Println(s)
}

func sortRecursive[T interface{}](s []T, start int, end int, less func(x, y T) bool) {
	if start >= end-1 {
		return
	}
	partition(s, start, end, less)
}

func partition[T interface{}](s []T, start int, end int, Less func(x, y T) bool) {
	p, pval := end-1, s[end-1]

	for i := start; i < end-1; i++ {
		if Less(s[i], pval) {
			s[p], s[i] = s[i], s[p]
			p++
		}
	}
}
