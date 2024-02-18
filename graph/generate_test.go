package graph

import (
	"math/rand"
	"testing"
)

func TestPruefer(t *testing.T) {
	rng := rand.New(rand.NewSource(RAND_SEED))
	const nverts = 1000
	pruefer := make([]int, nverts-2)
	for i := 0; i < len(pruefer); i++ {
		pruefer[i] = rng.Intn(nverts)
	}
	tree, e := NewTreeFromPruefer(pruefer)
	if e != nil {
		t.Errorf("could not generate Tree from valid Pruefer sequence %v", pruefer)
	}
	prueferRecovered := PrueferFromTree(tree)
	if len(prueferRecovered) != len(pruefer) {
		t.Errorf("recovered Pruefer sequence has incorrect length %v (should be %v)", len(prueferRecovered), len(pruefer))
	}
	for i := 0; i < len(pruefer); i++ {
		if pruefer[i] != prueferRecovered[i] {
			t.Errorf("recovered Pruefer sequence has incorrect value %v at index %v (should be %v)", prueferRecovered[i], i, pruefer[i])
		}
	}
}
