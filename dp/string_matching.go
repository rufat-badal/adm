package dp

import "fmt"

type StringCompareCase int

const (
	Match StringCompareCase = iota
	Insert
	Delete
	None
)

type Cell struct {
	Cost   int
	Parent StringCompareCase
}

func CompareStrings(first, second string) (int, [][]Cell) {
	const CostInsert = 1
	const CostDelete = 1
	const CostSubstitute = 1

	firstRunes := []rune(first)
	secondRunes := []rune(second)

	m := make([][]Cell, len(firstRunes)+1)
	for i := range m {
		m[i] = make([]Cell, len(secondRunes)+1)
	}

	m[0][0].Cost = 0
	m[0][0].Parent = None

	for i := 1; i <= len(firstRunes); i++ {
		m[i][0].Cost = i * CostDelete
		m[i][0].Parent = Delete
	}

	for j := 1; j <= len(secondRunes); j++ {
		m[0][j].Cost = j * CostInsert
		m[0][j].Parent = Insert
	}

	opt := make([]int, 3)
	var matchc int
	var fr, sr rune
	for i := 1; i <= len(firstRunes); i++ {
		for j := 1; j <= len(secondRunes); j++ {
			fr = firstRunes[i-1]
			sr = secondRunes[j-1]
			if fr != sr {
				matchc = CostSubstitute
			} else {
				matchc = 0
			}
			opt[Match] = m[i-1][j-1].Cost + matchc
			opt[Insert] = m[i][j-1].Cost + CostInsert
			opt[Delete] = m[i-1][j].Cost + CostDelete

			m[i][j].Cost = opt[Match]
			m[i][j].Parent = Match
			for k := Insert; k <= Delete; k++ {
				if opt[k] < m[i][j].Cost {
					m[i][j].Cost = opt[k]
					m[i][j].Parent = k
				}
			}
		}
	}

	return m[len(firstRunes)][len(secondRunes)].Cost, m
}

func reconstructPathCompareStringsRecursive(first, second []rune, i, j int, m [][]Cell) string {
	if m[i][j].Parent == None {
		return ""
	}

	var partialOut string
	if m[i][j].Parent == Match {
		if first[i-1] == second[j-1] {
			partialOut = " M"
		} else {
			partialOut = fmt.Sprintf(" S[%s -> %s]", string(first[i-1]), string(second[j-1]))
		}
		return reconstructPathCompareStringsRecursive(first, second, i-1, j-1, m) + partialOut
	}
	if m[i][j].Parent == Insert {
		partialOut = fmt.Sprintf(" I[%s]", string(second[j-1]))
		return reconstructPathCompareStringsRecursive(first, second, i, j-1, m) + partialOut
	}
	// m[i][j].Parent must be Delete
	partialOut = fmt.Sprintf(" D[%s]", string(first[i-1]))
	return reconstructPathCompareStringsRecursive(first, second, i-1, j, m) + partialOut
}

func ReconstructPathCompareStrings(first, second string, m [][]Cell) (string, error) {
	firstRunes := []rune(first)
	secondRunes := []rune(second)
	nrows := len(m)
	if nrows != len(firstRunes)+1 {
		return "", fmt.Errorf("m has %v rows (should have %v)", nrows, len(firstRunes)+1)
	}
	ncols := len(m[0])
	if ncols != len(secondRunes)+1 {
		return "", fmt.Errorf("m has %v cols (should have %v)", ncols, len(secondRunes)+1)
	}

	path := reconstructPathCompareStringsRecursive(firstRunes, secondRunes, len(firstRunes), len(secondRunes), m)

	return path[1:], nil
}
