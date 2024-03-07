package main

import (
	"fmt"

	"github.com/rufat-badal/adm/combinatorics"
)

func PrintSet(s []bool) {
	fmt.Print("{")
	first := 1
	for _, inSet := range s {
		if inSet {
			fmt.Printf("%v", first)
			break
		} else {
			first++
		}
	}
	if first <= len(s) {
		for i, inSet := range s[first:] {
			if inSet {
				fmt.Printf(", %v", first+i+1)
			}
		}
	}
	fmt.Print("}\n")
}

func main() {
	const n = 7
	nSets := 0
	for s := range combinatorics.GenerateAllSubsets(n) {
		PrintSet(s)
		nSets++
	}
	nSetsShould := 1
	for i := 0; i < n; i++ {
		nSetsShould *= 2
	}
	fmt.Printf("number of subsets: %v (should be %v)\n", nSets, nSetsShould)
}
