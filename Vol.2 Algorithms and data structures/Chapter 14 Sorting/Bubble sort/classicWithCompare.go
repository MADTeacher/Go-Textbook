package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func BubleSort(arr []int, comp Compare) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if comp(arr[j], arr[j+1]) {
				swap(j, j+1)
			}
		}
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := BubleSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = BubleSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
