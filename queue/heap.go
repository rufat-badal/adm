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
	SetWeight(int, int)
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

type heapContainer[T interface{}] struct {
	data []HeapItem[T]
}

func (hc heapContainer[T]) Get(i int) HeapItem[T] {
	return hc.data[i]
}

func (hc heapContainer[T]) Swap(i, j int) {
	hc.data[i], hc.data[j] = hc.data[j], hc.data[i]
}

func (hc heapContainer[T]) Len() int {
	return len(hc.data)
}

func (hc heapContainer[T]) Append(it HeapItem[T]) heapContainerInterface[T] {
	return heapContainer[T]{append(hc.data, it)}
}

func (hc heapContainer[T]) Pop() heapContainerInterface[T] {
	return heapContainer[T]{hc.data[:len(hc.data)-1]}
}

func (hc heapContainer[T]) SetWeight(i, newWeight int) {
	hc.data[i].Weight = newWeight
}

type MinHeap[T interface{}] struct {
	cnt heapContainerInterface[T]
}

func newMinHeapFromContainer[T interface{}](cnt heapContainerInterface[T]) MinHeap[T] {
	for i := cnt.Len()/2 - 1; i >= 0; i-- {
		bubbleDown[T](cnt, i)
	}
	return MinHeap[T]{cnt}
}

func NewMinHeap[T interface{}](items []HeapItem[T]) MinHeap[T] {
	hc := heapContainer[T]{items}
	return newMinHeapFromContainer[T](hc)
}

func (h MinHeap[T]) Len() int {
	return h.cnt.Len()
}

func (h *MinHeap[T]) Insert(val T, w int) {
	h.cnt = h.cnt.Append(HeapItem[T]{val, w})
	bubbleUp(h.cnt, h.cnt.Len()-1)
}

func (h MinHeap[T]) String() string {
	stringData := make([]string, h.cnt.Len())
	for i := 0; i < h.cnt.Len(); i++ {
		stringData[i] = fmt.Sprintf("%v", h.cnt.Get(i))
	}
	return fmt.Sprintf("%v", stringData)
}

func (h MinHeap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MinHeap[T]) ExtractMin() (HeapItem[T], error) {
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

type heapContainerByValueAccess[T comparable] struct {
	data    []HeapItem[T]
	indexOf map[T]int
}

func (hc heapContainerByValueAccess[T]) Get(i int) HeapItem[T] {
	return hc.data[i]
}

func (hc heapContainerByValueAccess[T]) Swap(i, j int) {
	x, y := hc.data[i], hc.data[j]
	hc.data[i], hc.data[j] = y, x
	hc.indexOf[x.Value], hc.indexOf[y.Value] = j, i
}

func (hc heapContainerByValueAccess[T]) Len() int {
	return len(hc.data)
}

func (hc heapContainerByValueAccess[T]) Append(it HeapItem[T]) heapContainerInterface[T] {
	newIndexOf := hc.indexOf
	newIndexOf[it.Value] = hc.Len()
	newData := append(hc.data, it)
	return heapContainerByValueAccess[T]{newData, newIndexOf}
}

func (hc heapContainerByValueAccess[T]) Pop() heapContainerInterface[T] {
	last := hc.data[len(hc.data)-1]
	delete(hc.indexOf, last.Value)
	return heapContainerByValueAccess[T]{hc.data[:len(hc.data)-1], hc.indexOf}
}

func (hc heapContainerByValueAccess[T]) SetWeight(i, newWeight int) {
	hc.data[i].Weight = newWeight
}

type MinHeapWithDecreaseWeight[T comparable] struct {
	hs      MinHeap[T]
	indexOf map[T]int
}

func NewMinHeapWithDecreaseWeight[T comparable](items []HeapItem[T]) MinHeapWithDecreaseWeight[T] {
	// No copy of items is made!
	indexOf := make(map[T]int)
	for i, it := range items {
		indexOf[it.Value] = i
	}
	hc := heapContainerByValueAccess[T]{items, indexOf}
	hs := newMinHeapFromContainer[T](hc)
	return MinHeapWithDecreaseWeight[T]{hs, hc.indexOf}
}

func (h MinHeapWithDecreaseWeight[T]) Len() int {
	return h.hs.Len()
}

func (h *MinHeapWithDecreaseWeight[T]) Insert(val T, w int) {
	h.hs.Insert(val, w)
}

func (h MinHeapWithDecreaseWeight[T]) String() string {
	return h.hs.String()
}

func (h MinHeapWithDecreaseWeight[T]) IsEmpty() bool {
	return h.hs.IsEmpty()
}

func (h *MinHeapWithDecreaseWeight[T]) ExtractMin() (HeapItem[T], error) {
	return h.hs.ExtractMin()
}

func (h MinHeapWithDecreaseWeight[T]) DecreaseWeight(value T, newWeight int) int {
	// DecreaseWeight does nothing if no item with value 'value' is present or the newWeight is larger or equal
	i, ok := h.indexOf[value]
	if !ok {
		return -1
	}
	it := h.hs.cnt.Get(i)
	if it.Weight <= newWeight {
		return it.Weight
	}
	oldWeight := it.Weight
	h.hs.cnt.SetWeight(i, newWeight)
	bubbleUp(h.hs.cnt, i)
	return oldWeight
}
