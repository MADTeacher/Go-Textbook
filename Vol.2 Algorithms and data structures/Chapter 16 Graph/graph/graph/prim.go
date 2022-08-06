package graph

import (
	"fmt"
	s "graph/collections"
)

type primNode[T Vertex] struct {
	To     T
	From   T
	weight int
}

func (d *primNode[T]) GetKey() int {
	return d.weight
}

func (d *primNode[T]) GetName() string {
	return fmt.Sprintf("%v", d.From)
}

func (g *Graph[T]) Prim(start T) *Graph[T] {
	set := s.NewSet[T]()
	mst := NewNotDirecteGraph[T]()

	set.Add(start)
	for set.Size() != g.AmountVertex() {
		heap := s.NewHeap[*primNode[T]](g.AmountEdges())
		set.ForEach(func(vertex *T) {
			g.ForEachAdjacentEdge(*vertex, func(edge *AdjacentEdge[T]) {
				if set.Contains(edge.To) {
					return
				}
				heap.Insert(&primNode[T]{
					From:   *vertex,
					To:     edge.To,
					weight: edge.Weight,
				})
			})
		})

		for !heap.IsEmpty() {
			v, _ := heap.Remove()
			mst.AddEdge(v.From, v.To, v.weight)
			set.Add(v.To)
			break
		}
	}

	return mst
}
