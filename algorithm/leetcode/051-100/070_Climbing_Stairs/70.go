package main

import "fmt"

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	first, second := 2, 3
	for i := 5; i <= n; i++ {
		current := first + second
		first = second
		second = current
	}
	return first + second
}

func climbStairs3(n int) int {
	if n <= 3 {
		return n
	}
	dp := make(map[int]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func climbStairs1(n int) int {
	if n <= 3 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

func climbStairsOptimize(n int) int {
	if n <= 3 {
		return n
	}
	first, second := 2, 3
	for i := 5; i <= n; i++ {
		current := first + second
		first = second
		second = current
	}
	return first + second
}

func main() {
	fmt.Println(climbStairsOptimize(1) == 1)
	fmt.Println(climbStairsOptimize(2) == 2)
	fmt.Println(climbStairsOptimize(3) == 3)
	fmt.Println(climbStairsOptimize(4) == 5)
	fmt.Println(climbStairsOptimize(5) == 8)
	fmt.Println(climbStairsOptimize(6) == 13)
	fmt.Println(climbStairsOptimize(7) == 21)
	fmt.Println(climbStairsOptimize(8) == 34)
	fmt.Println(climbStairsOptimize(9) == 55)
	fmt.Println(climbStairsOptimize(10) == 89)
	fmt.Println(climbStairsOptimize(11) == 144)
	fmt.Println(climbStairsOptimize(41) == 267914296)
	fmt.Println(climbStairsOptimize(35) == 14930352)
}
