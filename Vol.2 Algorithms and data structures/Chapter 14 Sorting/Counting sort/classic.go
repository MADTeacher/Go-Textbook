package main

import (
	"errors"
	"fmt"
	"math"
)

func findMaxMin(arr []int) (aMax, aMin int, err error) {
	if len(arr) == 0 {
		return 0, 0, errors.New("array is empty")
	}

	aMax = math.MinInt
	aMin = math.MaxInt
	for _, v := range arr {
		if aMax < v {
			aMax = v
		}
		if aMin > v {
			aMin = v
		}
	}
	return
}

func CountingSort(arr []int) ([]int, error) {
	aMax, aMin, err := findMaxMin(arr)
	if err != nil {
		return nil, err
	}
	arrayCounts := make([]int, aMax-aMin+1)

	for _, it := range arr {
		arrayCounts[it-aMin]++
	}

	it := 0
	for idx, count := range arrayCounts {
		for count > 0 {
			arr[it] = idx + aMin
			it++
			count--
		}
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, -2, -4, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	CountingSort(arr)
	fmt.Printf("Array after sorting: %v\n", arr)
}
