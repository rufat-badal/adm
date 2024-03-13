package graph

import (
	"github.com/rufat-badal/adm/queue"
)

type TSPSolution struct {
	Solution []int
	Cost     int
}

func (s TSPSolution) Extend(e Edge) TSPSolution {
	sol := make([]int, len(s.Solution)+1)
	copy(sol, s.Solution)
	sol[len(sol)-1] = e.Head
	return TSPSolution{sol, s.Cost + e.Weight}
}

func tspCostLowerBound(psol TSPSolution, min int, g Graph) int {
	return psol.Cost + min*(g.NumVertices-len(psol.Solution)+1)
}

func minEdgeWeight(g Graph) int {
	min := MAXINT
	for _, edges := range g.Edges {
		for _, edge := range edges {
			if edge.Weight < min {
				min = edge.Weight
			}
		}
	}
	return min
}

func isSolution(ps TSPSolution, g Graph) bool {
	return len(ps.Solution) == g.NumVertices
}

func TravelingSalesman(g Graph) TSPSolution {
	// TSP always starts at 1
	// We assume that there are no duplicate edges between two vertices
	min := minEdgeWeight(g)
	firstPsol := TSPSolution{[]int{1}, 0}
	psols := []queue.HeapItem[TSPSolution]{{Value: firstPsol, Weight: tspCostLowerBound(firstPsol, min, g)}}
	q := queue.NewMinHeap[TSPSolution](psols)
	minCost := MAXINT
	var sol TSPSolution

	for hit, e := q.ExtractMin(); e == nil; hit, e = q.ExtractMin() {
		ps := hit.Value
		if ps.Cost >= minCost {
			break
		}
		if isSolution(ps, g) && ps.Cost < minCost {
			sol = ps
			minCost = ps.Cost
		}
		i := ps.Solution[len(ps.Solution)-1]
	edgeLoop:
		for _, edge := range g.Edges[i] {
			for _, j := range ps.Solution {
				if j == edge.Head {
					continue edgeLoop
				}
			}
			psNew := ps.Extend(edge)
			q.Insert(psNew, tspCostLowerBound(psNew, min, g))
		}
	}

	return sol
}
