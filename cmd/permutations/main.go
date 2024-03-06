package main

import (
	"fmt"

	"github.com/rufat-badal/adm/combinatorics"
)

func main() {
	const n = 5
	numPerms := 0
	for p := range combinatorics.GenerateAllPermutations(n) {
		fmt.Println(p)
		numPerms++
	}
	numPermsShould := 1
	for i := 2; i <= n; i++ {
		numPermsShould *= i
	}
	fmt.Printf("number of permutations: %v (should be %v)\n", numPerms, numPermsShould)
}
