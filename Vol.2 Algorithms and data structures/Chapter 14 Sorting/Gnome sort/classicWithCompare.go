package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func GnomeSort(arr []int, comp Compare) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	i := 1
	for i < len(arr) {
		if comp(arr[i], arr[i-1]) {
			i++
		} else {
			swap(i, i-1)
			if i > 1 {
				i--
			}
		}
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := GnomeSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = GnomeSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
