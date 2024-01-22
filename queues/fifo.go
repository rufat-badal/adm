package queue

import (
	"errors"
	"fmt"
)

type FIFOQueueItemType interface{}

type FIFOQueue[T FIFOQueueItemType] struct {
	data   []T
	start  int
	Length int
}

const startLength = 4

func NewFIFOQueue[T FIFOQueueItemType]() FIFOQueue[T] {
	return FIFOQueue[T]{make([]T, startLength), 0, 0}
}

func (q *FIFOQueue[T]) increaseCapacity() {
	c := len(q.data)
	if q.Length == c {
		newData := make([]T, 2*c)
		for i := 0; i < q.Length; i++ {
			newData[i] = q.data[(q.start+i)%c]
		}
		q.data = newData
		q.start = 0
	}
}

func (q FIFOQueue[T]) String() string {
	stringData := make([]string, q.Length)
	for i := 0; i < q.Length; i++ {
		stringData[i] = fmt.Sprintf("%v", q.data[(q.start+i)%len(q.data)])
	}
	return fmt.Sprintf("%v", stringData)
}

func (q *FIFOQueue[T]) Capacity() int {
	return len(q.data)
}

func (q *FIFOQueue[T]) Enqueue(item T) {
	q.increaseCapacity()
	q.data[(q.start+q.Length)%len(q.data)] = item
	q.Length++
}

func (q *FIFOQueue[T]) decreaseCapacity() {
	c := len(q.data)
	if q.Length < c/4 {
		newData := make([]T, c/2)
		for i := 0; i < q.Length; i++ {
			newData[i] = q.data[(q.start+i)%c]
		}
		q.data = newData
		q.start = 0
	}
}

func (q *FIFOQueue[T]) Dequeue() (T, error) {
	if q.Length == 0 {
		return *new(T), errors.New("Dequeue called on empty queue")
	}

	first := q.data[q.start]
	q.start = (q.start + 1) % len(q.data)
	q.Length--
	q.decreaseCapacity()
	return first, nil
}

func (q FIFOQueue[T]) IsEmpty() bool {
	return q.Length == 0
}
