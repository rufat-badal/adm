package graph

import (
	"math/rand"
	"testing"
)

type sortableItem struct {
	Id     int
	Weight int
}

func TestSort(t *testing.T) {
	rng := rand.New(rand.NewSource(RAND_SEED))
	const nitems = 1000000
	s := make([]sortableItem, nitems)
	for i := 0; i < len(s); i++ {
		s[i] = sortableItem{rng.Int(), i}
	}
	rng.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	Sort[sortableItem](s, func(x, y sortableItem) bool { return x.Weight < y.Weight })
	for i, it := range s {
		if it.Weight != i {
			t.Errorf("item with wrong weight %v at index %v, expected weight %v", it.Weight, i, i)
		}
	}
}
