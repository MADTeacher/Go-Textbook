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

func Knapsack(knapsackСapacity int, items []Item) int {
	dp := make([]int, knapsackСapacity+1)
	for i := 1; i < len(items)+1; i++ {
		for w := knapsackСapacity; w >= 0; w-- {
			if items[i-1].Weight <= w {
				item := items[i-1]
				dp[w] = max(dp[w], dp[w-item.Weight]+item.Cost)
			}
		}
	}
	return dp[knapsackСapacity]
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
	result := Knapsack(knapsackСapacity, items)
	fmt.Printf("MaxCost: %d", result)
}
