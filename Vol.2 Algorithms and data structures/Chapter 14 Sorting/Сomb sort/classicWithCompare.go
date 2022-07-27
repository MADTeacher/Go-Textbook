package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func СombSort(arr []int, comp Compare) ([]int, error) {
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
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := СombSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = СombSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
