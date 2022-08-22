package main

import (
	"fmt"
	"math"
)

func findMinRoute(tsp [][]int, startCity int) (int, []int) {
	sum := 0
	counter := 0
	i, j := 0, startCity
	min := math.MaxInt
	visitedCity := make(map[int]struct{})

	visitedCity[startCity] = struct{}{}
	route := make([]int, len(tsp))

	for i < len(tsp) && j < len(tsp[i]) {
		if counter >= len(tsp[i])-1 {
			break
		}

		_, ok := visitedCity[j]
		if j != i && !ok {
			if tsp[i][j] < min {
				min = tsp[i][j]
				route[counter] = j + 1
			}
		}
		j++

		if j == len(tsp[i]) {
			sum += min
			min = math.MaxInt
			visitedCity[route[counter]-1] = struct{}{}
			j = 0
			i = route[counter] - 1
			counter++
		}
	}

	i = route[counter-1] - 1
	sum += tsp[i][startCity]
	route[counter] = startCity + 1

	route = append([]int{startCity + 1}, route...)
	fmt.Println(route)
	return sum, route
}

func main() {
	tsp := [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}
	profit, route := findMinRoute(tsp, 1)
	fmt.Printf("Profit: %d, route: %v\n", profit, route)
}
