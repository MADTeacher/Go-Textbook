package graph

import (
	"fmt"
	s "graph/collections"
)

type Vertex interface {
	comparable
}

type Edge[T Vertex] struct {
	From   T
	To     T
	Weight int
}

type AdjacentEdge[T Vertex] struct {
	To     T
	Weight int
}

type Graph[T Vertex] struct {
	Vertexes      map[T]*s.Set[AdjacentEdge[T]]
	IsNotDirected bool
}

func NewNotDirecteGraph[T Vertex]() *Graph[T] {
	vertexes := map[T]*s.Set[AdjacentEdge[T]]{}
	return &Graph[T]{
		Vertexes:      vertexes,
		IsNotDirected: true,
	}
}

func NewDirecteGraph[T Vertex]() *Graph[T] {
	vertexes := map[T]*s.Set[AdjacentEdge[T]]{}
	return &Graph[T]{
		Vertexes: vertexes,
	}
}

func (g *Graph[T]) AddVertex(vertex T) {
	_, ok := g.Vertexes[vertex]
	if !ok {
		g.Vertexes[vertex] = s.NewSet[AdjacentEdge[T]]()
	}
}

func (g *Graph[T]) AddEdge(vertex1, vertex2 T, weight int) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)

	g.Vertexes[vertex1].Add(AdjacentEdge[T]{vertex2, weight})
	if g.IsNotDirected {
		g.Vertexes[vertex2].Add(AdjacentEdge[T]{vertex1, weight})
	}
}

func (g *Graph[T]) AddEdgeWithoutWeight(vertex1, vertex2 T) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)

	g.Vertexes[vertex1].Add(AdjacentEdge[T]{vertex2, 0})
	if g.IsNotDirected {
		g.Vertexes[vertex2].Add(AdjacentEdge[T]{vertex1, 0})
	}
}

func (g *Graph[T]) ForEachVertex(f func(vertex *T)) {
	for v, _ := range g.Vertexes {
		f(&v)
	}
}

func (g *Graph[T]) ForEachEdge(f func(edge *Edge[T])) {
	for vertex1, edges := range g.Vertexes {
		edges.ForEach(func(value *AdjacentEdge[T]) {
			f(&Edge[T]{vertex1, value.To, value.Weight})
		})
	}
}

func (g *Graph[T]) ForEachAdjacentEdge(vertex T, f func(edge *AdjacentEdge[T])) {
	edges, ok := g.Vertexes[vertex]
	if ok {
		edges.ForEach(func(value *AdjacentEdge[T]) {
			f(value)
		})
	}
}

func (g *Graph[T]) AmountEdges() int {
	count := 0

	for _, edges := range g.Vertexes {
		count += edges.Size()
	}
	if g.IsNotDirected {
		count /= 2
	}
	return count
}

func (g *Graph[T]) AmountVertex() int {
	return len(g.Vertexes)
}

func (g *Graph[T]) PrintAllEdges() {
	g.ForEachEdge(func(edge *Edge[T]) {
		fmt.Printf("Edge(V1: %v, V2: %v, Weight: %v)\n",
			edge.From, edge.To, edge.Weight)
	})
}

func (g *Graph[T]) PrintAllVertexes() {
	g.ForEachVertex(func(vertex *T) {
		fmt.Printf("Vertex: %v\n",
			*vertex)
	})
}
