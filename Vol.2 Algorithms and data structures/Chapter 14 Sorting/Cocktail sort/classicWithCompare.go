package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func CocktailSort(arr []int, comp Compare) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		for i := right; i > left; i-- {
			if comp(arr[i-1], arr[i]) {
				swap(i-1, i)
			}
		}
		left++
		for i := left; i < right; i++ {
			if comp(arr[i], arr[i+1]) {
				swap(i, i+1)
			}
		}
		right--
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := CocktailSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = CocktailSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
