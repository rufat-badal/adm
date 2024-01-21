package main

import (
	"fmt"

	queue "github.com/rufat-badal/adm/queues"
)

func main() {
	q := queue.NewFIFOQueue[int]()
	for i := 0; i < 17; i++ {
		q.Enqueue(i)
		fmt.Printf("Enqueued %v, queue: %v, queue capacity: %v\n", i, q, q.Capacity())
	}
	for i := 0; i < 13; i++ {
		x, e := q.Dequeue()
		if e != nil {
			break
		}
		fmt.Printf("Dequeued %v, queue: %v, queue capacity: %v\n", x, q, q.Capacity())
	}
	for i := 17; i < 33; i++ {
		q.Enqueue(i)
		fmt.Printf("Enqueued %v, queue: %v, queue capacity: %v\n", i, q, q.Capacity())
	}
}
