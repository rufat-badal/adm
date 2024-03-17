package dp

func BinomialCoefficient(n, k int) int {
	bc := make([][]int, 2)
	bc[0] = make([]int, n+1)
	bc[1] = make([]int, n+1)
	bc[0][0] = 1
	bc[1][0] = 1
	old := 0
	new := 1

	for m := 1; m < n; m++ {
		bc[new][m] = 1
		for l := 1; l < m; l++ {
			bc[new][l] = bc[old][l-1] + bc[old][l]
		}
		new, old = (new+1)%2, new
	}

	bc[new][n] = 1
	for l := 1; l <= k; l++ {
		bc[new][l] = bc[old][l-1] + bc[old][l]
	}
	return bc[new][k]
}
