package graph

import "math/rand"

func Sort[T interface{}](s []T, isSmaller func(x T, y T) bool) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}
