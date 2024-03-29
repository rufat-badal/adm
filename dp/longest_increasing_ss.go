package dp

import "fmt"

func LongestIncreasingSubsequence(seq []int) (int, []int) {
	// l[i] = "length of the longest increasing subsequence of seq ending at i"
	l := make([]int, len(seq))
	l[0] = 1
	parent := make([]int, len(seq))
	parent[0] = -1

	var ml, p int
	for i := 1; i < len(l); i++ {
		ml = 0
		p = -1
		for j := 0; j < i; j++ {
			if seq[j] > seq[i] {
				continue
			}
			if l[j] > ml {
				ml = l[j]
				p = j
			}
		}
		l[i] = ml + 1
		parent[i] = p
	}

	return l[len(l)-1], parent
}

func reconstructLIS(i int, seq []int, parent []int) []int {
	if i == 0 {
		return seq[:1]
	}

	return append(reconstructLIS(parent[i], seq, parent), seq[i])
}

func ReconstructLIS(seq []int, parent []int) ([]int, error) {
	var ss []int
	if len(seq) != len(parent) {
		return ss, fmt.Errorf("length of seq (%v) does not coincide with length of parent (%v)", len(seq), len(parent))
	}
	for i, p := range parent {
		if p < -1 || p >= len(seq) {
			return ss, fmt.Errorf("invalid parent slice provided, p[%v] = %v, but it must be in 0...%v", i, p, len(seq)-1)
		}
	}

	return reconstructLIS(len(seq)-1, seq, parent), nil
}
