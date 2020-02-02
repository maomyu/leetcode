package main

import "fmt"

// 旋转数组的最小值
func getmin(data []int) int {
	j := len(data) - 1
	for j > 0 {
		for j > 0 && data[j-1] < data[j] {
			j--
		}
		if j > 0 && data[j-1] > data[j] {
			return data[j]
		}
		if j == 0 {
			break
		}
	}
	return data[0]
}

// 双指针方法
func GetMin(data []int) int {
	if len(data) == 0 {
		return 0
	}
	i, j := 0, len(data)-1
	for i < j {
		if data[i] > data[j] {
			i++
		} else {
			break
		}
	}
	return data[i]
}
func main() {
	data := []int{1, 1, 3, 4, 5}
	// data1 := []int{5}
	fmt.Println(GetMin(data))
}
