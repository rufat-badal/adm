package main

import (
	"fmt"

	"github.com/rufat-badal/adm/dp"
)

func main() {
	first := "Rufxyzatち"
	second := "Rufat"
	cost, m := dp.CompareStrings(first, second)
	fmt.Println(cost)
	s, e := dp.ReconstructPathCompareStrings(first, second, m)
	if e == nil {
		fmt.Println(s)
	}
}
