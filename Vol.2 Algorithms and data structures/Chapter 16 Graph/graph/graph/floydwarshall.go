package graph

func (g *Graph[T]) FloydWarshall() map[T]map[T]int {
	costMatrix := make(map[T]map[T]int)

	g.ForEachVertex(func(vertex *T) {
		costMatrix[*vertex] = map[T]int{}
		g.ForEachAdjacentEdge(*vertex, func(edge *AdjacentEdge[T]) {
			costMatrix[*vertex][edge.To] = edge.Weight
		})
	})
	g.ForEachVertex(func(k *T) {
		g.ForEachVertex(func(u *T) {
			g.ForEachVertex(func(v *T) {
				_, ok1 := costMatrix[*u][*k]
				_, ok2 := costMatrix[*k][*v]
				if !ok1 || !ok2 {
					return
				}

				oldCost, ok := costMatrix[*u][*v]
				newCost := costMatrix[*u][*k] + costMatrix[*k][*v]
				if ok {
					if newCost < oldCost {
						costMatrix[*u][*v] = newCost
					}
				} else {
					costMatrix[*u][*v] = newCost
				}
			})
		})
	})
	return costMatrix
}
