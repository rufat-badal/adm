package main

import (
	"fmt"

	"github.com/rufat-badal/adm/dp"
)

func main() {
	seq := []int{1, 7, 3, 5, 6, 8, 100, 9, 13, 15}
	l, p := dp.LongestIncreasingSubsequence(seq)
	fmt.Println(l)
	fmt.Println(dp.ReconstructLIS(seq, p))
}
