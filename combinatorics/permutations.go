package combinatorics

func generatePerumationsBacktrack(p []int, k int, n int, c chan []int) {
	if k == n {
		pCopy := make([]int, n)
		copy(pCopy, p)
		c <- pCopy
		lastPerm := true
		for i, a := range p {
			if a != n-i {
				lastPerm = false
				break
			}
		}
		if lastPerm {
			close(c)
		}
		return
	}

	inPerm := make([]bool, n)
	for i := 0; i < k; i++ {
		inPerm[p[i]-1] = true
	}
	nc := 0
	candidates := make([]int, n)
	for i, contained := range inPerm {
		if !contained {
			candidates[nc] = i + 1
			nc++
		}
	}
	k++
	for i := 0; i < nc; i++ {
		p[k-1] = candidates[i]
		generatePerumationsBacktrack(p, k, n, c)
	}
}

func GenerateAllPermutations(n int) <-chan []int {
	c := make(chan []int)
	go func() {
		p := make([]int, n)
		generatePerumationsBacktrack(p, 0, n, c)
	}()
	return c
}
