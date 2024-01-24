package queue

import (
	"errors"
	"fmt"
)

type StackItemType interface{}

type Stack[T StackItemType] struct {
	data []T
}

func NewStack[T StackItemType]() Stack[T] {
	return Stack[T]{make([]T, 0)}
}

func (s Stack[T]) String() string {
	stringData := make([]string, len(s.data))
	for i := 0; i < len(s.data); i++ {
		stringData[i] = fmt.Sprintf("%v", s.data[i])
	}
	return fmt.Sprintf("%v", stringData)
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) decreaseCapacity() {
	if len(s.data) < cap(s.data)/4 {
		newData := make([]T, len(s.data), cap(s.data)/2)
		copy(newData, s.data)
		s.data = newData
	}
}

func (s Stack[T]) Length() int {
	return len(s.data)
}

func (s Stack[T]) Capacity() int {
	return cap(s.data)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return *new(T), errors.New("cannot pop an element from an empty stack")
	}
	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.decreaseCapacity()
	return popped, nil
}
