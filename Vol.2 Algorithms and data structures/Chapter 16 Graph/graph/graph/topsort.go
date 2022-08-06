package graph

import "errors"

func (g *Graph[T]) TopologicalSort() ([]T, map[T]int, error) {
	inDegree := map[T]int{}
	g.ForEachEdge(func(edge *Edge[T]) {
		_, ok := inDegree[edge.From]
		if !ok {
			inDegree[edge.From] = 0
		}

		_, ok = inDegree[edge.To]
		if ok {
			inDegree[edge.To] += 1
		} else {
			inDegree[edge.To] = 1
		}
	})

	resultList := []T{}
	resultMap := map[T]int{}
	countLevel := 0
	for len(inDegree) > 0 {
		tempList := []T{}
		for vertex, deegre := range inDegree {
			if deegre == 0 {
				tempList = append(tempList, vertex)
				resultMap[vertex] = countLevel
			}
		}
		if len(tempList) == 0 {
			return nil, nil, errors.New("found loop")
		}
		for _, vertex := range tempList {
			g.ForEachAdjacentEdge(vertex, func(edge *AdjacentEdge[T]) {
				inDegree[edge.To]--
			})
			delete(inDegree, vertex)
			resultList = append(resultList, vertex)
		}
		countLevel++
	}
	return resultList, resultMap, nil
}
