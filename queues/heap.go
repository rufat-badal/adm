package queue

import (
	"errors"
	"fmt"
)

type HeapItem[T comparable] struct {
	Value  T
	Weight int
}

type MinHeap[T comparable] struct {
	data []HeapItem[T]
	// All items in the queue must have distinct values!
	indexOf map[T]int
}

func NewMinHeap[T comparable]() MinHeap[T] {
	return MinHeap[T]{make([]HeapItem[T], 0), make(map[T]int)}
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

func (h *MinHeap[T]) swap(i, j int) {
	x, y := h.data[i], h.data[j]
	h.data[i], h.data[j] = y, x
	h.indexOf[x.Value], h.indexOf[y.Value] = j, i
}

func (h *MinHeap[T]) bubbleUp(id int) {
	pid := Parent(id)
	if pid == -1 {
		return
	}
	p, c := h.data[pid], h.data[id]
	if p.Weight <= c.Weight {
		return
	}
	h.swap(pid, id)
	h.bubbleUp(pid)
}

func (h *MinHeap[T]) Insert(item HeapItem[T]) error {
	_, valuePresent := h.indexOf[item.Value]
	if valuePresent {
		return fmt.Errorf("an item with the same value %v was already inerted into the queue", item.Value)
	}
	h.data = append(h.data, item)
	id := len(h.data) - 1
	h.indexOf[item.Value] = id
	h.bubbleUp(id)
	return nil
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

func (h *MinHeap[T]) bubbleDown(id int) {
	min := id
	cid := FirstChild(id)
	// Find index of the node of minimal weight in the family of the node at p
	for i := 0; i < 2; i++ {
		if (cid + i) >= len(h.data) {
			break
		}
		if h.data[cid+i].Weight < h.data[min].Weight {
			min = cid + i
		}
	}
	if min == id {
		return
	}
	h.swap(id, min)
	h.bubbleDown(min)
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
	delete(h.indexOf, min.Value)
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.indexOf[last.Value] = 0
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 1 {
		h.bubbleDown(0)
	}
	h.decreaseCapacity()
	return min, nil
}
