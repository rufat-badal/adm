package main

import (
	"fmt"

	"github.com/rufat-badal/adm/dp"
)

func main() {
	cost, _ := dp.CompareStrings("Rufat", "Ruufat")
	fmt.Println(cost)
}
