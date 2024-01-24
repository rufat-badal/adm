package main

import (
	"fmt"
	"math/rand"

	queue "github.com/rufat-badal/adm/queues"
)

func main() {
	h := queue.NewMinHeap[int]()
	for i := 0; i < 10; i++ {
		h.Insert(queue.HeapItem[int]{Value: i, Weight: rand.Intn(100)})
		fmt.Printf("Inserted an item into the heap, data: %v, length: %v, capacity: %v\n", h, h.Length(), h.Capacity())
	}
	for i := 0; i < 10; i++ {
		min, _ := h.ExtractMin()
		fmt.Printf("Extracted minimum %v from heap, data: %v, length: %v, capacity: %v\n", min, h, h.Length(), h.Capacity())
	}
}
