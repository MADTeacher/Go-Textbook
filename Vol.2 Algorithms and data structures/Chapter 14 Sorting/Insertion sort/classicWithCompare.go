package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func InsertionSort(arr []int, comp Compare) ([]int, error) {
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
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := InsertionSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = InsertionSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
