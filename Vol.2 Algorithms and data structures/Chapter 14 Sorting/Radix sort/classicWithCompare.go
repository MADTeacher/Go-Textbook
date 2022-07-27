package main

import (
	"errors"
	"fmt"
	"math"
)

func getMax(arr []int) int {
	aMax := math.MinInt
	for _, val := range arr {
		if aMax < val {
			aMax = val
		}
	}
	return aMax
}

func RadixSort(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	digPlace := 1
	result := make([]int, len(arr))

	maxNumber := getMax(arr)
	for maxNumber/digPlace > 0 {
		count := make([]int, 10) // частота чисел от 0 до 9
		for _, val := range arr {
			count[(val/digPlace)%10]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			result[count[(arr[i]/digPlace)%10]-1] = arr[i]
			count[(arr[i]/digPlace)%10]--
		}
		copy(arr, result)
		digPlace *= 10
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 6, 0, 362, 214, 22, 54, 109, 5, 3}
	fmt.Printf("Array before sort: %v\n", arr)
	sortedArray, _ := RadixSort(arr)
	fmt.Printf("Array after sorting: %v\n", sortedArray)
}
