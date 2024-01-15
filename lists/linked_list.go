package lists

import "fmt"

func Hello() {
	fmt.Println("Hello from linked_list")
}

type LinkedList[T comparable] struct {
	Next  *LinkedList[T]
	Value T
}

func LinkedListFromSlice[T comparable](s []T) *LinkedList[T] {
	if len(s) == 0 {
		return nil
	}

	l := &LinkedList[T]{nil, s[len(s)-1]}
	for i := len(s) - 2; i >= 0; i-- {
		new_head := LinkedList[T]{l, s[i]}
		l = &new_head
	}
	return l
}
