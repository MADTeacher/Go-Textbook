package main

import (
	"errors"
	"fmt"
)

type Compare func(a, b int) bool

func mergeSortImpl(arr, buffer []int, l, r int, comp Compare) {
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

func MergeSort(arr []int, comp Compare) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	buffer := make([]int, len(arr))
	mergeSortImpl(arr, buffer, 0, len(arr)-1, comp)
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := MergeSort(arr, func(a, b int) bool {
		return a < b // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = MergeSort(arr, func(a, b int) bool {
		return a > b // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
