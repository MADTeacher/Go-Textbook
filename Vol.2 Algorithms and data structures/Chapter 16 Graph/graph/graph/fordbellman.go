package graph

import "math"

type node[T Vertex] struct {
	cost        int
	predecessor *T
}

func (g *Graph[T]) FordBellman(start, end T) (*[]T, int) {
	nodes := map[T]*node[T]{}

	g.ForEachVertex(func(vertex *T) {
		nodes[*vertex] = &node[T]{math.MaxInt16, nil}
	})
	nodes[start].cost = 0

	amountVertex := g.AmountVertex()
	for i := 0; i < amountVertex-1; i++ {
		g.ForEachEdge(func(edge *Edge[T]) {
			cost := nodes[edge.From].cost + edge.Weight
			if cost < nodes[edge.To].cost {
				nodes[edge.To].cost = cost
				vertex := edge.From
				nodes[edge.To].predecessor = &vertex
			}
		})
	}

	// проверка на наличие отрицательного цикла
	hasNegativeLoop := false
	g.ForEachEdge(func(edge *Edge[T]) {
		if nodes[edge.From].cost+edge.Weight <
			nodes[edge.To].cost {
			hasNegativeLoop = true
		}
	})

	if hasNegativeLoop {
		return nil, 0
	}

	path := []T{}
	for vertex := &end; vertex != nil; vertex = nodes[*vertex].predecessor {
		path = append([]T{*vertex}, path...)
	}
	return &path, nodes[end].cost
}
