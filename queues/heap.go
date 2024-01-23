package queue

type HeapValueType interface{}

type HeapItem[T HeapValueType] struct {
	Value  T
	Weight int
}

type MinHeap[T HeapValueType] struct {
	data   []HeapItem[T]
	Length int
}

func NewMinHeap[T HeapValueType]() MinHeap[T] {
	return MinHeap[T]{make([]HeapItem[T], 1), 0}
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
