package combinatorics

func generateSubsetsBacktrack(s []bool, k int, n int, c chan []bool) {
	if k == n {
		sInverted := make([]bool, n)
		// This assures the order {}, {1}, {2}, {1, 2}, ...
		for i, inSet := range s {
			sInverted[len(s)-1-i] = inSet
		}
		c <- sInverted
		lastSet := true
		for _, inSet := range s {
			if !inSet {
				lastSet = false
				break
			}
		}
		if lastSet {
			close(c)
		}
		return
	}

	candidates := make([]bool, 2)
	candidates[1] = true
	k++
	for _, cand := range candidates {
		s[k-1] = cand
		generateSubsetsBacktrack(s, k, n, c)
	}
}

func GenerateAllSubsets(n int) <-chan []bool {
	c := make(chan []bool)
	go func() {
		s := make([]bool, n)
		generateSubsetsBacktrack(s, 0, n, c)
	}()
	return c
}
