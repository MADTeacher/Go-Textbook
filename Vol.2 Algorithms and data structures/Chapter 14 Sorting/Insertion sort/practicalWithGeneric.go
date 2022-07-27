package main

import (
	"errors"
	"fmt"
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

///////////////////////////

type Compare[T any] func(a, b T) bool

func InsertionSort[T any](arr []T, comp Compare[T]) ([]T, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		it := i
		for ; it > 0 && comp(arr[it-1], temp); it-- {
			arr[it] = arr[it-1]
		}
		arr[it] = temp
	}
	return arr, nil
}

func main() {
	workerSlice := []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4},
		{"George", 3}, {"Max", 60}, {"Tommy", 94}, {"William", 12},
		{"Sophia", 14}, {"Oliver", 13}, {"Sandra", 91},
		{"Ann", 6}, {"Elizabeth", 9}, {"Kate", 20}}

	fmt.Printf("Array before sort: %v\n", workerSlice)
	fmt.Println("---------Sort by id-----------")
	sortedArray, _ := InsertionSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() < b.GetID() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = InsertionSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() > b.GetID() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
	fmt.Println("---------Sort by name-----------")
	sortedArray, _ = InsertionSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() < b.GetName() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = InsertionSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() > b.GetName() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
