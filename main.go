package main

import (
	"fmt"

	"github.com/rufat-badal/adm/cmd/dp"
)

func main() {
	for n := 0; n < 20; n++ {
		fmt.Println(dp.Fibonacci(n))
	}
}
