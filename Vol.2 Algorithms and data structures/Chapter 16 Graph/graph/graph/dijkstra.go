package graph

import (
	"fmt"
	s "graph/collections"
	"math"
)

type dijkstraNode[T Vertex] struct {
	vertex      T
	distance    int
	predecessor *dijkstraNode[T]
}

func (d *dijkstraNode[T]) GetKey() int {
	return d.distance
}

func (d *dijkstraNode[T]) GetName() string {
	return fmt.Sprintf("%v", d.vertex)
}

func (g *Graph[T]) Dijkstra(start, end T) (*[]T, int) {
	// инициализируем структуры данных
	heap := s.NewHeap[*dijkstraNode[T]](g.AmountVertex())
	nodes := map[T]*dijkstraNode[T]{}
	g.ForEachVertex(func(vertex *T) {
		node := &dijkstraNode[T]{
			vertex:   *vertex,
			distance: math.MaxInt,
		}
		heap.Insert(node)
		nodes[*vertex] = node
	})

	// устанавливаем значение пути начальной вершины в 0
	nodes[start].distance = 0
	heap.Change(nodes[start])

	// реализуем обход и расчет пути для каждой из вершины, начиная с
	// вершины start до тех пор, пока не дойдем
	// до вершины end
	for !heap.IsEmpty() {
		v, _ := heap.Remove()
		g.ForEachAdjacentEdge(v.vertex, func(edge *AdjacentEdge[T]) {
			node := nodes[edge.To]

			if node == nil {
				return
			}

			if v.distance+edge.Weight < node.distance {
				node.distance = v.distance + edge.Weight
				node.predecessor = v
				heap.Change(node)
			}
		})

		if v.vertex == end {
			path := []T{}
			cost := v.distance
			for it := v; it != nil; it = it.predecessor {
				path = append([]T{it.vertex}, path...)
			}
			return &path, cost
		}
	}
	return nil, 0
}
