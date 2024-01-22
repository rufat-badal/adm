package queue

import (
	"errors"
	"fmt"
)

type StackItemType interface{}

type Stack[T StackItemType] struct {
	data   []T
	Length int
}

func NewStack[T StackItemType]() Stack[T] {
	return Stack[T]{make([]T, 0), 0}
}

func (s Stack[T]) String() string {
	stringData := make([]string, s.Length)
	for i := 0; i < s.Length; i++ {
		stringData[i] = fmt.Sprintf("%v", s.data[i])
	}
	return fmt.Sprintf("%v", stringData)
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
	s.Length++
}

func (s *Stack[T]) decreaseCapacity() {
	if s.Length < cap(s.data)/4 {
		newData := make([]T, cap(s.data)/2)
		copy(newData, s.data)
		s.data = newData
	}
}

func (s Stack[T]) Capacity() int {
	return cap(s.data)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Length == 0 {
		return *new(T), errors.New("cannot pop an element from an empty stack")
	}
	popped := s.data[s.Length-1]
	s.data = s.data[:s.Length-1]
	s.Length--
	s.decreaseCapacity()
	return popped, nil
}
