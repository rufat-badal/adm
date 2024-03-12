package queue

import (
	"math/rand"
	"testing"
)

func sortedHeapItems(nitems int) []HeapItem[int] {
	items := make([]HeapItem[int], nitems)
	for i := 0; i < nitems; i++ {
		items[i] = HeapItem[int]{i, i}
	}
	return items
}

func shuffledItems[T interface{}](items []HeapItem[T], rng *rand.Rand) []HeapItem[T] {
	itemsShuffled := make([]HeapItem[T], len(items))
	copy(itemsShuffled, items)
	rng.Shuffle(len(items), func(i, j int) { itemsShuffled[i], itemsShuffled[j] = itemsShuffled[j], itemsShuffled[i] })
	return itemsShuffled
}

func TestInsert(t *testing.T) {
	rng := rand.New(rand.NewSource(RAND_SEED))
	const nitems = 1000
	items := sortedHeapItems(nitems)
	h := NewMinHeap[int](shuffledItems[int](items, rng))
	if h.cnt.Len() != len(items) {
		t.Errorf("heap was not created correctly, it contains %v items but should contain %v", h.cnt.Len(), len(items))
	}
	for _, it := range items {
		itemFound := false
		for j := 0; j < h.cnt.Len(); j++ {
			if h.cnt.Get(j) == it {
				itemFound = true
			}
		}
		if !itemFound {
			t.Errorf("item %v was not found in the heap", it)
		}
	}
}

type extractMiner[T interface{}] interface {
	ExtractMin() (HeapItem[T], error)
}

func testExtractAll(t *testing.T, h extractMiner[int], nitems int) {
	for i := 0; i < nitems; i++ {
		x, e := h.ExtractMin()
		if e != nil {
			t.Errorf("Expected item %v, but queue returned <nil>", HeapItem[int]{i, i})
		}
		if x.Value != i {
			t.Errorf("Expected item %v, but got item %v", HeapItem[int]{i, i}, x)
		}
	}
}

func TestExtractMin(t *testing.T) {
	rng := rand.New(rand.NewSource(RAND_SEED))
	const nitems = 1000
	items := sortedHeapItems(nitems)
	h := NewMinHeap[int](shuffledItems[int](items, rng))
	testExtractAll(t, &h, nitems)
}

type decreaseOp[T comparable] struct {
	Value T
	Delta int
}

func TestDecreaseWeight(t *testing.T) {
	rng := rand.New(rand.NewSource(RAND_SEED))
	const nitems = 1000
	const ndecreases = 100
	const maxDecrease = 100
	items := sortedHeapItems(nitems)
	var decreases [ndecreases]decreaseOp[int]
	for i := 0; i < ndecreases; i++ {
		decreases[i] = decreaseOp[int]{rng.Intn(nitems), rng.Intn(maxDecrease) + 1}
	}
	for _, d := range decreases {
		items[d.Value].Weight += d.Delta
	}
	h := NewMinHeapWithDecreaseWeight[int](shuffledItems[int](items, rng))
	for _, d := range decreases {
		h.DecreaseWeight(d.Value, items[d.Value].Weight-d.Delta)
		items[d.Value].Weight -= d.Delta
	}
	testExtractAll(t, &h, nitems)
}
