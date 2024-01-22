package queue

type HeapItem[T interface{}] struct {
	Value  T
	Weight int
}
