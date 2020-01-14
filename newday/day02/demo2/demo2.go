package main

import (
	"fmt"
)

// 整数反转
func reverse(x int) int {
	min := -1 << 31
	max := 1<<31 - 1
	sign := 0
	if x < 0 {
		sign = -1
		x = -x
	}
	sum := 0
	for x != 0 {
		temp := x % 10
		sum = sum*10 + temp
		x = x / 10
	}
	if sign == 0 {
		if sum < max {
			return sum
		}
	} else {
		if -sum > min {
			return -sum
		}
	}

	return 0
}

func main() {
	fmt.Println(reverse(1234569999999))
}
