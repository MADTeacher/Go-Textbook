package main

import (
	"fmt"
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func QuickSelect(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	pivot := arr[generator.Intn(len(arr))]
	var L, M, R []int
	for _, val := range arr {
		if pivot > val {
			L = append(L, val)
		}
		if pivot == val {
			M = append(M, val)
		}
		if pivot < val {
			R = append(R, val)
		}
	}
	if k <= len(L) {
		return QuickSelect(L, k)
	} else if k <= (len(L) + len(M)) {
		return pivot
	} else {
		return QuickSelect(R, k-(len(L)+len(M)))
	}
}

func main() {
	arr := []int{3, -2, 0, 4, 22, -1, 34, 10, 5, 7, 9}

	fmt.Printf("Array: %v\n", arr)
	k := 3
	kMin := QuickSelect(arr, k)
	fmt.Printf("%v-th min element is: %v\n", k, kMin)
	k = 2
	kMin = QuickSelect(arr, k)
	fmt.Printf("%v-th min element is: %v\n", k, kMin)
	k = 5
	kMin = QuickSelect(arr, k)
	fmt.Printf("%v-th min element is: %v\n", k, kMin)
}
