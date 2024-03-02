package graph

type ResidualEdge struct {
	Head     int
	Flow     int
	Residual int
}

type VertexPair struct {
	Tail int
	Head int
}

type ResidualGraph struct {
	Edges            [][]ResidualEdge
	vertexPairToEdge map[VertexPair]*ResidualEdge
	NumVertices      int
}

func newEmptyResidualGraph(nvertices int) ResidualGraph {
	return ResidualGraph{make([][]ResidualEdge, nvertices), make(map[VertexPair]*ResidualEdge), nvertices}
}

func (g Graph) ResidualGraph() ResidualGraph {
	// edge weights are interpreted as capacities
	rg := newEmptyResidualGraph(g.NumVertices)
	for tail, edges := range g.Edges {
		for _, e := range edges {
			head := e.Head
			capacity := e.Weight
			forwardPair := VertexPair{tail, head}
			re, ok := rg.vertexPairToEdge[forwardPair]
			if ok {
				// an edge between tail and head was already encountered
				re.Residual += capacity
			} else {
				// add forward residual edge
				rg.Edges[tail] = append(rg.Edges[tail], ResidualEdge{Head: head, Flow: 0, Residual: capacity})
				rg.vertexPairToEdge[forwardPair] = &rg.Edges[tail][len(rg.Edges[tail])-1]
				// add backward residual edge
				rg.Edges[head] = append(rg.Edges[head], ResidualEdge{Head: tail, Flow: 0, Residual: 0})
				backwardPair := VertexPair{head, tail}
				rg.vertexPairToEdge[backwardPair] = &rg.Edges[head][len(rg.Edges[head])-1]
			}
		}
	}

	return rg
}
