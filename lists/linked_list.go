package lists

import "fmt"

func Hello() {
	fmt.Println("Hello from linked_list")
}

type LinkedList[T comparable] struct {
	Next  *LinkedList[T]
	Value T
}

func Insert[T comparable](l *LinkedList[T], x T) *LinkedList[T] {
	return &LinkedList[T]{l, x}
}

func LinkedListFromSlice[T comparable](s []T) *LinkedList[T] {
	if len(s) == 0 {
		return nil
	}

	l := &LinkedList[T]{nil, s[len(s)-1]}
	for i := len(s) - 2; i >= 0; i-- {
		l = Insert[T](l, s[i])
	}
	return l
}
