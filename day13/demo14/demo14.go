package main

import "fmt"

// 求x的平方根
func mySqrt(x int) int {
	left := 0
	right := x/2 + 1
	for left <= right {
		mid := (right + left) / 2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}
func main() {
	fmt.Println(mySqrt(1))
}
