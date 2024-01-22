package main

import (
	"fmt"

	queue "github.com/rufat-badal/adm/queues"
)

func main() {
	s := queue.NewStack[int]()
	fmt.Println(s)
	for i := 0; i < 20; i++ {
		s.Push(i + 42)
		fmt.Printf("Pushed %v on stack, stack: %v, length: %v, capacity: %v\n", i+42, s, s.Length, s.Capacity())
	}
	for i := 0; i < 15; i++ {
		popped, _ := s.Pop()
		fmt.Printf("Popped %v of the stack, stack: %v, capacity: %v\n", popped, s, s.Capacity())
	}
}
