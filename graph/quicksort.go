package graph

import (
	"math/rand"
)

func Sort[T interface{}](s []T, less func(x, y T) bool) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	sortRecursive[T](s, 0, len(s), less)
}

func sortRecursive[T interface{}](s []T, start int, end int, less func(x, y T) bool) {
	if start >= end-1 {
		return
	}
	mid := partition(s, start, end, less)
	sortRecursive[T](s, start, mid, less)
	sortRecursive[T](s, mid+1, end, less)
}

func partition[T interface{}](s []T, start int, end int, less func(x, y T) bool) int {
	pval := s[end-1]
	p := start

	for i := start; i < end-1; i++ {
		if less(s[i], pval) {
			s[p], s[i] = s[i], s[p]
			p++
		}
	}

	s[p], s[end-1] = pval, s[p]

	return p
}
