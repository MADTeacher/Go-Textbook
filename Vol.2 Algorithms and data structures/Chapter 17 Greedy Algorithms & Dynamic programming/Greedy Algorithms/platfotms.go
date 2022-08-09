package main

import (
	"fmt"
	"sort"
)

type Train struct {
	Arrival   float64
	Departure float64
	ID        int
}

// Для сортировки по времени прибытия
type TrainArrival []Train

func (t TrainArrival) Len() int {
	return len(t)
}

func (t TrainArrival) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TrainArrival) Less(i, j int) bool {
	return t[i].Arrival < t[j].Arrival
}

// Для сортировки по времени отправления
type TrainDeparture []Train

func (t TrainDeparture) Len() int {
	return len(t)
}

func (t TrainDeparture) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TrainDeparture) Less(i, j int) bool {
	return t[i].Departure < t[j].Departure
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMinPlatforms(trains []Train) int {
	platforms, count := 0, 0
	arrival := trains
	departure := make([]Train, len(trains))
	copy(departure, trains)

	sort.Sort(TrainArrival(arrival))
	sort.Sort(TrainDeparture(departure))

	i, j := 0, 0
	for i < len(trains) {
		if arrival[i].Arrival < departure[j].Departure {
			count++
			platforms = max(count, platforms)
			// переходим к рассмотрению времени следующего пребывающего поезда
			i++
		} else {
			count--
			// переходим к рассмотрению времени отправления следующего поезда
			j++
		}
	}
	return platforms
}

func main() {
	trains := []Train{
		{2.00, 2.30, 1},
		{2.10, 3.40, 10},
		{3.00, 3.20, 5},
		{3.20, 4.30, 6},
		{3.15, 4.00, 7},
		{4.50, 5.10, 2},
	}
	result := FindMinPlatforms(trains)
	fmt.Printf("Platforms %d: ", result)
}
