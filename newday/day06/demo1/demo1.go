package main

import "fmt"

// 爬楼梯
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairs_two(n int) int {
	f0, f1 := 1, 1
	if n == 0 || n == 1 {
		return f0
	}
	f2 := 0
	for i := 2; i <= n; i++ {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
	}
	return f2
}
func main() {
	fmt.Println(climbStairs_two(5))
}
