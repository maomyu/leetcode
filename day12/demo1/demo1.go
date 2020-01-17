package main

import "fmt"

// 求最大次方,如果n为2的倍数
func getMax(n int) int {
	if n == 2 {
		return 1
	}
	return getMax(n/2) + 1
}
func getMax_two(n int) int {
	if n <= 1 {
		return 0
	}

	sum := 2
	cout := 1
	for true {
		sum = sum * 2
		if sum <= n {
			cout++
		} else {
			break
		}
	}
	return cout
}
func main() {
	fmt.Println(getMax_two(63))
}
