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

func СombSort[T any](arr []T, comp Compare[T]) ([]T, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	const factor float64 = 1.247
	var step float64 = float64(len(arr) - 1)

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for step >= 1 {
		for i := 0; float64(i)+step < float64(len(arr)); i++ {
			if comp(arr[i], arr[i+int(step)]) {
				swap(i, i+int(step))
			}
		}
		step /= factor
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
	sortedArray, _ := СombSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() < b.GetID() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = СombSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() > b.GetID() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
	fmt.Println("---------Sort by name-----------")
	sortedArray, _ = СombSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() < b.GetName() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = СombSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() > b.GetName() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
