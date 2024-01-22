package main

import (
	"fmt"

	queue "github.com/rufat-badal/adm/queues"
)

func main() {
	q := queue.NewFIFOQueue[int]()
	fmt.Println(q)
	for i := 0; i < 20; i++ {
		q.Enqueue(i + 42)
		fmt.Printf("Enqueued %v, queue: %v, length: %v, capacity: %v\n", i+42, q, q.Length, q.Capacity())
	}
	for i := 0; i < 15; i++ {
		popped, _ := q.Dequeue()
		fmt.Printf("Dequeued %v, queue: %v, capacity: %v\n", popped, q, q.Capacity())
	}
}
