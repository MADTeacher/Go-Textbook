package main

import (
	"errors"
	"fmt"
	"sort"
)

////// Worker /////////
type Worker struct {
	Name string
	Id   uint8
}

func (c *Worker) GetID() int {
	return int(c.Id)
}

func (c *Worker) GetName() string {
	return c.Name
}

////////////поиск по id/////////////
func InterpolationSearch(arr []Worker, x uint8) (int, error) {
	if x < uint8(arr[0].GetID()) || x > uint8(arr[len(arr)-1].GetID()) {
		return -1, errors.New("id not found")
	}
	high := len(arr) - 1
	low := 0
	for arr[high].GetID() != arr[low].GetID() &&
		(x >= uint8(arr[low].GetID()) && x <= uint8(arr[high].GetID())) {
		pos := low + int(float64(high-low)/
			float64(arr[high].GetID()-arr[low].GetID())*
			float64(x-uint8(arr[low].GetID())))
		if uint8(arr[pos].GetID()) == x {
			return pos, nil
		} else if uint8(arr[pos].GetID()) < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1, errors.New("id not found")
}

func main() {
	workerSlice := []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4},
		{"George", 3}, {"Max", 60}, {"Tommy", 94}, {"William", 12},
		{"Sophia", 14}, {"Oliver", 13}, {"Sandra", 91},
		{"Ann", 6}, {"Elizabeth", 9}, {"Kate", 20}}
	// сортировка по возрастанию id
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetID() < workerSlice[j].GetID()
	})
	fmt.Printf("Array after sorting by id: %v\n", workerSlice)
	index, _ := InterpolationSearch(workerSlice, 13) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := InterpolationSearch(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
}
