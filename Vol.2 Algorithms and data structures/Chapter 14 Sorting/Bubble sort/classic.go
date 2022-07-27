package main

import (
	"errors"
	"fmt"
)

func BubleSort(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				swap(j, j+1)
			}
		}
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	sortedArray, _ := BubleSort(arr)
	fmt.Printf("Array before sort: %v\n", arr)
	fmt.Printf("Array after sort: %v\n", sortedArray)
}
