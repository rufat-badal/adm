package queue

import (
	"math/rand"
	"testing"
)

func randomMinHeap[T comparable](items []HeapItem[T]) MinHeap[T] {
	h := NewMinHeap[T]()
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	for _, x := range items {
		h.Insert(x)
	}
	return h
}

func TestExtractMin(t *testing.T) {
	const nitems = 1000
	items := make([]HeapItem[int], nitems)
	for i := 0; i < nitems; i++ {
		items[i] = HeapItem[int]{i, i}
	}
	h := randomMinHeap[int](items)
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
