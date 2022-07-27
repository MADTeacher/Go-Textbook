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

func mergeSortImpl[T any](arr, buffer []T, l, r int, comp Compare[T]) {
	if l < r {
		m := (l + r) / 2
		mergeSortImpl(arr, buffer, l, m, comp)
		mergeSortImpl(arr, buffer, m+1, r, comp)

		k, j := l, m+1
		for i := l; i <= m || j <= r; {
			if j > r || (i <= m && comp(arr[i], arr[j])) {
				buffer[k] = arr[i]
				i++
			} else {
				buffer[k] = arr[j]
				j++
			}
			k++
		}
		for i := l; i <= r; i++ {
			arr[i] = buffer[i]
		}
	}
}

func MergeSort[T any](arr []T, comp Compare[T]) ([]T, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	buffer := make([]T, len(arr))
	mergeSortImpl(arr, buffer, 0, len(arr)-1, comp)
	return arr, nil
}

func main() {
	workerSlice := []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4},
		{"George", 3}, {"Max", 60}, {"Tommy", 94}, {"William", 12},
		{"Sophia", 14}, {"Oliver", 13}, {"Sandra", 91},
		{"Ann", 6}, {"Elizabeth", 9}, {"Kate", 20}}

	fmt.Printf("Array before sort: %v\n", workerSlice)
	fmt.Println("---------Sort by id-----------")
	sortedArray, _ := MergeSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() > b.GetID() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = MergeSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() < b.GetID() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
	fmt.Println("---------Sort by name-----------")
	sortedArray, _ = MergeSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() > b.GetName() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = MergeSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() < b.GetName() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
