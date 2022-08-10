package main

import (
	"fmt"
)

func main() {
	array := []int{0, 2, 1, 6, 9, 10, 3, 2, 5, 6, 8, 15, 20, 0, 5}
	dp := make([]int, len(array))
	max, tempMax := 0, 0

	dp[0] = 1
	for i := 1; i < len(array); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if array[i] > array[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
			if dp[i] > tempMax {
				tempMax++
			} else {
				if tempMax > max {
					max = tempMax
				}
				tempMax = 0
			}
		}
	}
	fmt.Println(dp, max)
}
