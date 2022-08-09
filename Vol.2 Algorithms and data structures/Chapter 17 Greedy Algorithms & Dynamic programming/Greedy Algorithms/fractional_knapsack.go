package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Cost   int
	Weight int
	ID     int
}

type IncreaseItem []Item

func (items IncreaseItem) Len() int {
	return len(items)
}

func (items IncreaseItem) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items IncreaseItem) Less(i, j int) bool {
	a := float64(items[i].Cost) / float64(items[i].Weight)
	b := float64(items[j].Cost) / float64(items[j].Weight)
	return a > b
}

func fractionalKnapsack(knapsackСapacity int,
	items []Item) (finalCost float64, knapsack []Item) {
	sort.Sort(IncreaseItem(items))

	for i := 0; i < len(items); i++ {
		if items[i].Weight <= knapsackСapacity {
			knapsackСapacity -= items[i].Weight
			finalCost += float64(items[i].Cost)
			knapsack = append(knapsack, items[i])
		} else {
			weight := (float64(knapsackСapacity) /
				float64(items[i].Weight))
			lastCost := float64(items[i].Cost) * weight
			finalCost += lastCost
			knapsack = append(knapsack, Item{
				int(lastCost), int(weight), items[i].ID,
			})
			break
		}
	}
	return finalCost, knapsack
}

func main() {
	knapsackСapacity := 50
	items := []Item{{60, 15, 1}, {100, 10, 10}, {300, 50, 5}}
	maxCost, knapsack := fractionalKnapsack(knapsackСapacity, items)
	fmt.Println(maxCost, knapsack)
}
