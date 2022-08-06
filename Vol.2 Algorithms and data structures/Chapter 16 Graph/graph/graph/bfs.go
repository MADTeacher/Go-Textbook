package graph

import s "graph/collections"

func (g *Graph[T]) BFS(start T, walkFunc func(vertex *T) bool) {
	queue, visited := s.NewQueue[T](), s.NewSet[T]()
	queue.Enqueue(start)

	v, err := queue.Dequeue()
	for ; err == nil; v, err = queue.Dequeue() {

		if walkFunc(&v) {
			return
		}

		visited.Add(v)

		g.ForEachAdjacentEdge(v, func(edge *AdjacentEdge[T]) {
			if !visited.Contains(edge.To) {
				queue.Enqueue(edge.To)
			}
		})
	}
}
