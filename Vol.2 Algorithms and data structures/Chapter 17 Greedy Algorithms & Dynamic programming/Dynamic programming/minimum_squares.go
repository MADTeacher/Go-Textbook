package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func FindMinSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(FindMinSquares(20))
	fmt.Println(FindMinSquares(10))
	fmt.Println(FindMinSquares(96))
	fmt.Println(FindMinSquares(100))
}
