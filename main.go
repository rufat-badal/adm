package main

import (
	"fmt"

	"github.com/rufat-badal/adm/combinatorics"
)

func main() {
	su := combinatorics.NewSudoku([...]int{
		6, -1, -1, 2, 8, -1, -1, -1, -1,
		-1, 4, -1, -1, 9, -1, -1, -1, 8,
		-1, 3, 8, -1, 5, 7, 4, 9, -1,
		-1, -1, -1, 6, 7, 2, -1, 3, -1,
		-1, 1, -1, 5, 4, -1, -1, -1, -1,
		-1, 5, -1, 8, 3, -1, 6, -1, -1,
		8, -1, -1, -1, 1, -1, 9, 6, 2,
		-1, -1, 4, 7, 6, -1, 1, -1, 5,
		-1, 6, 5, 9, -1, -1, 3, -1, 4,
	})
	fmt.Println(su)
}
