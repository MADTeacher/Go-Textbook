package graph

import s "graph/collections"

func (g *Graph[T]) DFS(start T, walkFunc func(vertex *T) bool) {
	visited := s.NewSet[T]()
	isFound := false
	g.dfs(start, visited, &isFound, walkFunc)
}

func (g *Graph[T]) dfs(start T, visited *s.Set[T], isFound *bool, walkFunc func(vertex *T) bool) {
	visited.Add(start)

	*isFound = walkFunc(&start)
	if *isFound {
		return
	}

	g.ForEachAdjacentEdge(start, func(edge *AdjacentEdge[T]) {
		if *isFound {
			return
		}
		if !visited.Contains(edge.To) {
			g.dfs(edge.To, visited, isFound, walkFunc)
		}
	})
}
