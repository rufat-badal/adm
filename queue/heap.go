package queue

import (
	"errors"
	"fmt"
)

func Parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func FirstChild(i int) int {
	return 2*i + 1
}

type HeapItem[T interface{}] struct {
	Value  T
	Weight int
}

type heapContainerInterface[T interface{}] interface {
	Len() int
	Get(int) HeapItem[T]
	Swap(int, int)
	Append(it HeapItem[T]) heapContainerInterface[T]
	Pop() heapContainerInterface[T]
}

type heapContainerSimple[T interface{}] struct {
	data []HeapItem[T]
}

func (hc heapContainerSimple[T]) Get(i int) HeapItem[T] {
	return hc.data[i]
}

func (hc heapContainerSimple[T]) Swap(i, j int) {
	hc.data[i], hc.data[j] = hc.data[j], hc.data[i]
}

func (hc heapContainerSimple[T]) Len() int {
	return len(hc.data)
}

func (hc heapContainerSimple[T]) Append(it HeapItem[T]) heapContainerInterface[T] {
	return heapContainerSimple[T]{append(hc.data, it)}
}

func (hc heapContainerSimple[T]) Pop() heapContainerInterface[T] {
	return heapContainerSimple[T]{hc.data[:len(hc.data)-1]}
}

type MinHeapSimple[T interface{}] struct {
	cnt heapContainerInterface[T]
}

func bubbleDown[T interface{}](cnt heapContainerInterface[T], from int) {
	min := from
	cid := FirstChild(from)
	// Find index of the node of minimal weight in the family of the node at p
	for i := 0; i < 2; i++ {
		if (cid + i) >= cnt.Len() {
			break
		}
		if cnt.Get(cid+i).Weight < cnt.Get(min).Weight {
			min = cid + i
		}
	}
	if min == from {
		return
	}
	cnt.Swap(from, min)
	bubbleDown[T](cnt, min)
}

func bubbleUp[T interface{}](cnt heapContainerInterface[T], from int) {
	pid := Parent(from)
	if pid == -1 {
		return
	}
	p, c := cnt.Get(pid), cnt.Get(from)
	if p.Weight <= c.Weight {
		return
	}
	cnt.Swap(pid, from)
	bubbleUp(cnt, pid)
}

func newMinHeapFromContainer[T interface{}](cnt heapContainerInterface[T]) MinHeapSimple[T] {
	for i := cnt.Len()/2 - 1; i >= 0; i-- {
		bubbleDown[T](cnt, i)
	}
	return MinHeapSimple[T]{cnt}
}

func NewMinHeapSimple[T interface{}](items []HeapItem[T]) MinHeapSimple[T] {
	hc := heapContainerSimple[T]{items}
	return newMinHeapFromContainer[T](hc)
}

func (h MinHeapSimple[T]) Len() int {
	return h.cnt.Len()
}

func (h *MinHeapSimple[T]) Insert(it HeapItem[T]) {
	h.cnt = h.cnt.Append(it)
	bubbleUp(h.cnt, h.cnt.Len()-1)
}

func (h MinHeapSimple[T]) String() string {
	stringData := make([]string, h.cnt.Len())
	for i := 0; i < h.cnt.Len(); i++ {
		stringData[i] = fmt.Sprintf("%v", h.cnt.Get(i))
	}
	return fmt.Sprintf("%v", stringData)
}

func (h MinHeapSimple[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MinHeapSimple[T]) ExtractMin() (HeapItem[T], error) {
	if h.IsEmpty() {
		return *new(HeapItem[T]), errors.New("cannot extract minimum from an empty heap")
	}
	min := h.cnt.Get(0)
	h.cnt.Swap(0, h.cnt.Len()-1)
	h.cnt = h.cnt.Pop()
	if h.Len() > 1 {
		bubbleDown(h.cnt, 0)
	}
	return min, nil
}

type heapContainer[T comparable] struct {
	data    []HeapItem[T]
	indexOf map[T]int
}

func (hc heapContainer[T]) Get(i int) HeapItem[T] {
	return hc.data[i]
}

func (hc heapContainer[T]) Swap(i, j int) {
	x, y := hc.data[i], hc.data[j]
	hc.data[i], hc.data[j] = y, x
	hc.indexOf[x.Value], hc.indexOf[y.Value] = j, i
}

func (hc heapContainer[T]) Len() int {
	return len(hc.data)
}

func (hc heapContainer[T]) Append(it HeapItem[T]) heapContainerInterface[T] {
	newIndexOf := hc.indexOf
	newIndexOf[it.Value] = hc.Len()
	newData := append(hc.data, it)
	return heapContainer[T]{newData, newIndexOf}
}

func (hc heapContainer[T]) Pop() heapContainerInterface[T] {
	last := hc.data[len(hc.data)-1]
	delete(hc.indexOf, last.Value)
	return heapContainer[T]{hc.data[:len(hc.data)-1], hc.indexOf}
}

type MinHeapNew[T comparable] struct {
	h       MinHeapSimple[T]
	indexOf *map[T]int
}

func NewMinHeapNew[T comparable](items []HeapItem[T]) MinHeapNew[T] {
	// No copy of items is made!
	indexOf := make(map[T]int)
	for i, it := range items {
		indexOf[it.Value] = i
	}
	hc := heapContainer[T]{items, indexOf}
	h := newMinHeapFromContainer[T](hc)
	return MinHeapNew[T]{h, &hc.indexOf}
}

type MinHeap[T comparable] struct {
	data []HeapItem[T]
	// All items in the queue must have distinct values!
	indexOf map[T]int
}

func NewMinHeap[T comparable](items []HeapItem[T]) MinHeap[T] {
	// No copy of items is made!
	indexOf := make(map[T]int)
	for i, it := range items {
		indexOf[it.Value] = i
	}
	h := MinHeap[T]{items, indexOf}
	for i := len(items)/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
	return h
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

func (h *MinHeap[T]) DecreaseWeight(value T, newWeight int) int {
	// DecreaseWeight does nothing if no item with value 'value' is present or the newWeight is larger or equal
	id, ok := h.indexOf[value]
	if !ok {
		return -1
	}
	item := &h.data[id]
	if item.Weight <= newWeight {
		return item.Weight
	}
	oldWeight := item.Weight
	item.Weight = newWeight
	h.bubbleUp(id)
	return oldWeight
}
