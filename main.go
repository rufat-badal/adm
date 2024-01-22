package main

import (
	"fmt"

	queue "github.com/rufat-badal/adm/queues"
)

func main() {
	h := queue.NewMinHeap[int]()
	fmt.Println(h.Length)
}
