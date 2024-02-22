package graph

type UnionFind struct {
	nelems int
	nsets  int
	parent []int
	size   []int
	height []int
}

func NewUnionFind(nelems int) *UnionFind {
	parent := make([]int, nelems)
	for i := range parent {
		parent[i] = i
	}
	size := make([]int, nelems)
	for i := range size {
		size[i] = 1
	}
	height := make([]int, nelems)
	return &UnionFind{nelems, nelems, parent, size, height}
}

func (s *UnionFind) Find(elem int) int {
	if elem < 0 || elem >= s.nelems {
		return -1
	}

	if s.parent[elem] == elem {
		return elem
	} else {
		return s.Find(s.parent[elem])
	}
}

func (s *UnionFind) Union(elem1, elem2 int) {
	r1 := s.Find(elem1)
	if r1 == -1 {
		return
	}
	r2 := s.Find(elem2)
	if r2 == -1 || r1 == r2 {
		return
	}

	if s.size[r1] >= s.size[r2] {
		s.parent[r2] = r1
		s.size[r1] += s.size[r2]
		s.height[r1] = max(s.height[r1], s.height[r2]+1)
	} else {
		s.parent[r1] = r2
		s.size[r2] += s.size[r1]
		s.height[r2] = max(s.height[r2], s.height[r1]+1)
	}
	s.nsets--
}

func (s *UnionFind) SameComponent(i, j int) bool {
	return s.Find(i) == s.Find(j)
}

func (s *UnionFind) NumElems() int {
	return s.nelems
}

func (s *UnionFind) NumSets() int {
	return s.nsets
}

func (s *UnionFind) Height(elem int) int {
	r := s.Find(elem)
	if r == -1 {
		return -1
	}

	return s.height[r]
}
