package queue

import (
	"math/rand"
	"testing"
)

func randomMinHeap[T comparable](items []HeapItem[T]) MinHeap[T] {
	r := rand.New(rand.NewSource(RAND_SEED))
	itemsShuffled := make([]HeapItem[T], len(items))
	copy(itemsShuffled, items)
	r.Shuffle(len(items), func(i, j int) { itemsShuffled[i], itemsShuffled[j] = itemsShuffled[j], itemsShuffled[i] })
	return NewMinHeap[T](itemsShuffled)
}

func sortedHeapItems(nitems int) []HeapItem[int] {
	items := make([]HeapItem[int], nitems)
	for i := 0; i < nitems; i++ {
		items[i] = HeapItem[int]{i, i}
	}
	return items
}

func TestInsert(t *testing.T) {
	const nitems = 10000
	items := sortedHeapItems(nitems)
	h := randomMinHeap[int](items)
	for i := 0; i < nitems; i++ {
		it := h.data[h.indexOf[i]]
		if it.Weight != i {
			t.Errorf("item %v has wrong weight, should be %v", it, i)
		}
	}
}

func testExtractAll(t *testing.T, h MinHeap[int], nitems int) {
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
	const nitems = 10000
	items := sortedHeapItems(nitems)
	h := randomMinHeap[int](items)
	testExtractAll(t, h, nitems)
}

type decreaseOp[T comparable] struct {
	Value T
	Delta int
}

func TestDecreaseWeight(t *testing.T) {
	r := rand.New(rand.NewSource(RAND_SEED))
	const nitems = 10000
	const ndecreases = 1000
	const maxDecrease = 100
	items := sortedHeapItems(nitems)
	var decreases [ndecreases]decreaseOp[int]
	for i := 0; i < ndecreases; i++ {
		decreases[i] = decreaseOp[int]{r.Intn(nitems), r.Intn(maxDecrease) + 1}
	}
	for _, d := range decreases {
		items[d.Value].Weight += d.Delta
	}
	h := randomMinHeap[int](items)
	for _, d := range decreases {
		h.DecreaseWeight(d.Value, items[d.Value].Weight-d.Delta)
		items[d.Value].Weight -= d.Delta
	}
	testExtractAll(t, h, nitems)
}
