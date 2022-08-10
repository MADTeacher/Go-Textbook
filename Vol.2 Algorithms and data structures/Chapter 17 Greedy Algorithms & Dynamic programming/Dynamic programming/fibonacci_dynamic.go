package main

import (
	"fmt"
)

var dp [1000]int

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	if dp[n] > 2 {
		return dp[n]
	}

	k := 0
	if n&1 > 0 {
		k = (n + 1) / 2
	} else {
		k = n / 2
	}

	if n&1 > 0 {
		dp[n] = Fibonacci(k)*Fibonacci(k) + Fibonacci(k-1)*Fibonacci(k-1)
	} else {
		dp[n] = (2*Fibonacci(k-1) + Fibonacci(k)) * Fibonacci(k)
	}
	return dp[n]
}

func main() {
	fmt.Println(Fibonacci(20))
}
