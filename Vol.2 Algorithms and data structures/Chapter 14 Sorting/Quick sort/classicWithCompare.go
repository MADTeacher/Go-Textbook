package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func partition(arr []int, l, r int, comp Compare) int {
	x := arr[r]
	less := l

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := l; i < r; i++ {
		if comp(arr[i], x) {
			swap(i, less)
			less++
		}
	}
	swap(less, r)
	return less
}

func quickSortImpl(arr []int, l, r int, comp Compare) {
	if l < r {
		q := partition(arr, l, r, comp)
		quickSortImpl(arr, l, q-1, comp)
		quickSortImpl(arr, q+1, r, comp)
	}
}

func QuickSort(arr []int, comp Compare) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	quickSortImpl(arr, 0, len(arr)-1, comp)
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := QuickSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = QuickSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
