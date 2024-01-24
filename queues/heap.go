package queue

import (
	"errors"
	"fmt"
)

type HeapValueType interface{}

type HeapItem[T HeapValueType] struct {
	Value  T
	Weight int
}

type MinHeap[T HeapValueType] struct {
	data []HeapItem[T]
}

func NewMinHeap[T HeapValueType]() MinHeap[T] {
	return MinHeap[T]{make([]HeapItem[T], 0)}
}

func Parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func FirstChild(i int) int {
	return 2*i + 1
}

func (h MinHeap[T]) Length() int {
	return len(h.data)
}

func (h MinHeap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeap[T]) bubbleUp(idx int) {
	pIdx := Parent(idx)
	if pIdx == -1 {
		return
	}
	p := h.data[pIdx]
	c := h.data[idx]
	if p.Weight <= c.Weight {
		return
	}
	h.data[pIdx], h.data[idx] = c, p
	h.bubbleUp(pIdx)
}

func (h *MinHeap[T]) Insert(item HeapItem[T]) {
	h.data = append(h.data, item)
	h.bubbleUp(len(h.data) - 1)
}

func (h MinHeap[T]) Capacity() int {
	return cap(h.data)
}

func (h MinHeap[T]) String() string {
	stringData := make([]string, len(h.data))
	for i := 0; i < len(h.data); i++ {
		stringData[i] = fmt.Sprintf("%v", h.data[i])
	}
	return fmt.Sprintf("%v", stringData)
}

func (h *MinHeap[T]) bubbleDown(p int) {
	min := p
	c := FirstChild(p)
	// Find index of the node of minimal weight in the family of the node at p
	for i := 0; i < 2; i++ {
		if (c + i) >= len(h.data) {
			break
		}
		if h.data[c+i].Weight < h.data[min].Weight {
			min = c + i
		}
	}

	if min != p {
		h.data[p], h.data[min] = h.data[min], h.data[p]
		h.bubbleDown(min)
	}
}

func (h *MinHeap[T]) decreaseCapacity() {
	if len(h.data) < cap(h.data)/4 {
		newData := make([]HeapItem[T], len(h.data), cap(h.data)/2)
		copy(newData, h.data)
		h.data = newData
	}
}

func (h *MinHeap[T]) ExtractMin() (HeapItem[T], error) {
	if h.IsEmpty() {
		return *new(HeapItem[T]), errors.New("cannot extract minimum from an empty heap")
	}
	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.bubbleDown(0)
	}
	h.decreaseCapacity()
	return min, nil
}
