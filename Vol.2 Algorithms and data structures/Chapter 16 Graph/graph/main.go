package main

import (
	"fmt"
	"graph/graph"
)

// ///////////TopologicalSort/////////////
func main() {
	g := graph.NewDirecteGraph[string]()

	g.AddEdgeWithoutWeight("A", "B")
	g.AddEdgeWithoutWeight("A", "D")
	g.AddEdgeWithoutWeight("B", "C")
	g.AddEdgeWithoutWeight("B", "D")
	g.AddEdgeWithoutWeight("C", "E")
	g.AddEdgeWithoutWeight("D", "E")
	g.AddEdgeWithoutWeight("D", "F")
	g.AddEdgeWithoutWeight("E", "F")
	topList, topMap, _ := g.TopologicalSort()
	fmt.Println(topList)
	fmt.Println(topMap)
}

///////////////Prim///////////////////////
// func main() {
// 	vertexes := []graph.Edge[string]{
// 		{From: "A", To: "B", Weight: 13},
// 		{From: "A", To: "C", Weight: 6},
// 		{From: "A", To: "F", Weight: 4},
// 		{From: "B", To: "C", Weight: 7},
// 		{From: "B", To: "F", Weight: 7},
// 		{From: "B", To: "E", Weight: 5},
// 		{From: "C", To: "E", Weight: 1},
// 		{From: "C", To: "F", Weight: 8},
// 		{From: "E", To: "F", Weight: 9},
// 	}

// 	g := graph.NewNotDirecteGraph[string]()

// 	for _, it := range vertexes {
// 		g.AddEdge(it.From, it.To, it.Weight)
// 	}

// 	mst := g.Prim("C")
// 	fmt.Println(mst.AmountVertex() == g.AmountVertex())
// 	fmt.Println("---------Vertexes-----------")
// 	mst.PrintAllVertexes()
// 	fmt.Println("---------Edges-----------")
// 	mst.PrintAllEdges()
// }

///////////////FloydWarshall/////////////////////////
// func main() {
// 	vertexes := []graph.Edge[string]{
// 		{From: "A", To: "B", Weight: -3},
// 		{From: "B", To: "A", Weight: 4},
// 		{From: "B", To: "C", Weight: 5},
// 		{From: "B", To: "F", Weight: 7},
// 		{From: "C", To: "E", Weight: 1},
// 		{From: "C", To: "A", Weight: 6},
// 		{From: "E", To: "B", Weight: 5},
// 		{From: "E", To: "F", Weight: 6},
// 		{From: "F", To: "A", Weight: -4},
// 		{From: "F", To: "C", Weight: 8},
// 	}

// 	g := graph.NewDirecteGraph[string]()

// 	for _, it := range vertexes {
// 		g.AddEdge(it.From, it.To, it.Weight)
// 	}

// 	costMatrix := g.FloydWarshall()
// 	for key, val := range costMatrix {
// 		fmt.Printf("Key: %v, Val: %v\n", key, val)
// 	}
// }

///////////////FordBellman/////////////////////////
// func main() {
// 	vertexes := []graph.Edge[string]{
// 		{From: "A", To: "B", Weight: -3},
// 		{From: "B", To: "A", Weight: 4},
// 		{From: "B", To: "C", Weight: 5},
// 		{From: "B", To: "F", Weight: 7},
// 		{From: "C", To: "E", Weight: 1},
// 		{From: "C", To: "A", Weight: 6},
// 		{From: "E", To: "B", Weight: 5},
// 		{From: "E", To: "F", Weight: 6},
// 		{From: "F", To: "A", Weight: -4},
// 		{From: "F", To: "C", Weight: 8},
// 	}

// 	g := graph.NewDirecteGraph[string]()

// 	for _, it := range vertexes {
// 		g.AddEdge(it.From, it.To, it.Weight)
// 	}

// 	path, cost := g.FordBellman("E", "C")
// 	if path != nil {
// 		fmt.Printf("Path: %v with cost:%d\n", *path, cost)
// 	}
// 	path, cost = g.FordBellman("E", "A")
// 	if path != nil {
// 		fmt.Printf("Path: %v with cost:%d\n", *path, cost)
// 	}
// 	path, cost = g.FordBellman("A", "E")
// 	if path != nil {
// 		fmt.Printf("Path: %v with cost:%d\n", *path, cost)
// 	}
// }

///////////////// Dijkstra////////////////////
// func main() {
// 	vertexes := []graph.Edge[string]{
// 		{From: "A", To: "B", Weight: 8},
// 		{From: "A", To: "C", Weight: 5},
// 		{From: "B", To: "D", Weight: 1},
// 		{From: "B", To: "E", Weight: 13},
// 		{From: "C", To: "F", Weight: 14},
// 		{From: "C", To: "D", Weight: 10},
// 		{From: "F", To: "D", Weight: 9},
// 		{From: "F", To: "E", Weight: 6},
// 		{From: "D", To: "E", Weight: 8},
// 	}

// 	g := graph.NewNotDirecteGraph[string]()

// 	for _, it := range vertexes {
// 		g.AddEdge(it.From, it.To, it.Weight)
// 	}

// 	path, cost := g.Dijkstra("E", "A")
// 	fmt.Printf("Path: %v with cost:%d\n", *path, cost)
// }

//////////////////BSF-DSF///////////////////////////////////////////////////
// func main() {
// 	vertexes := []graph.Edge[string]{
// 		{From: "A", To: "B"},
// 		{From: "A", To: "C"},
// 		{From: "C", To: "F"},
// 		{From: "C", To: "G"},
// 		{From: "G", To: "M"},
// 		{From: "G", To: "N"},
// 		{From: "B", To: "D"},
// 		{From: "B", To: "E"},
// 		{From: "D", To: "H"},
// 		{From: "D", To: "I"},
// 		{From: "D", To: "J"},
// 		{From: "E", To: "K"},
// 		{From: "E", To: "L"},
// 	}

// 	g := graph.NewNotDirecteGraph[string]()

// 	for _, it := range vertexes {
// 		g.AddEdgeWithoutWeight(it.From, it.To)
// 	}

// 	path := []string{}
// 	g.DFS("A", func(vertex *string) bool {
// 		fmt.Printf("%s ", *vertex)
// 		path = append(path, *vertex)
// 		return *vertex == "K"
// 	})
// 	fmt.Println()

// 	path = Reverse(path)
// 	minPath := []string{"K"}
// 	find := "K"
// 	for _, v := range path {
// 		if g.Vertexes[v].Contains(graph.AdjacentEdge[string]{To: find}) {
// 			minPath = append(minPath, v)
// 			find = v
// 		}
// 	}
// 	fmt.Println(Reverse(minPath))
// }

// func Reverse[T any](input []T) []T {
// 	var output []T

// 	for i := len(input) - 1; i >= 0; i-- {
// 		output = append(output, input[i])
// 	}

// 	return output
// }

//////////////////base///////////////////////////////////////////////////
////// City /////////
// type City struct {
// 	Name string
// 	Id   uint8
// }

// func (c *City) GetID() int {
// 	return int(c.Id)
// }

// func (c *City) GetName() string {
// 	return c.Name
// }

// func main() {
// 	workerSlice := []City{{"Saint Petersburg", 2}, {"Moscow", 1}, {"Pskov", 3},
// 		{"Rostov-on-Don", 4}, {"Stavropol", 5}, {"Grozny", 7}, {"Gukovo", 6},
// 		{"Kalach-na-Donu", 14}, {"Kansk", 13}, {"Mamonovo", 91},
// 		{"Nizhnekamsk", 8}, {"Omsk", 9}, {"Oryol", 20}}

// 	g := graph.NewDirecteGraph[City]()

// 	for _, it := range workerSlice {
// 		g.AddVertex(it)
// 	}

// 	g.AddEdge(workerSlice[0], workerSlice[3], 1960)
// 	g.AddEdge(workerSlice[0], workerSlice[1], 1500)
// 	g.AddEdge(workerSlice[1], workerSlice[3], 460)
// 	g.AddEdge(workerSlice[3], workerSlice[6], 100)

// 	fmt.Println(len(workerSlice) == g.AmountVertex())
// 	fmt.Println("---------Vertexes-----------")
// 	g.PrintAllVertexes()
// 	fmt.Println("---------Edges-----------")
// 	g.PrintAllEdges()
// }
