package dp

type StringCompareCase int

const (
	Match StringCompareCase = iota
	Insert
	Delete
)

type Cell struct {
	Cost   int
	Parent StringCompareCase
}

func CompareStrings(first, second string) (int, [][]Cell) {
	lowestCost := max(len(first), len(second))
	m := make([][]Cell, len(first))
	for i := range m {
		m[i] = make([]Cell, len(second))
	}
	return lowestCost, m
}
