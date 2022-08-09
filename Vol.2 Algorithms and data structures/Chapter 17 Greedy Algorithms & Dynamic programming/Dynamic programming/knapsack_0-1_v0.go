package main

import "fmt"

type Item struct {
	Cost   int
	Weight int
	ID     int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Knapsack(knapsackСapacity, n int, items []Item) int {
	if n == 0 || knapsackСapacity == 0 {
		return 0
	}

	if items[n-1].Weight > knapsackСapacity {
		return Knapsack(knapsackСapacity, n-1, items)
	} else {
		return max(items[n-1].Cost+Knapsack(
			knapsackСapacity-items[n-1].Weight,
			n-1,
			items,
		), Knapsack(
			knapsackСapacity,
			n-1,
			items,
		))
	}
}

func main() {
	knapsackСapacity := 50
	items := []Item{
		{60, 10, 1},
		{100, 20, 10},
		{120, 30, 5},
		{100, 25, 6},
		{90, 19, 7},
	}
	result := Knapsack(knapsackСapacity, len(items), items)
	fmt.Printf("MaxCost: %d", result)
}
