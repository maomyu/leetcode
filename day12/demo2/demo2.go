package main

import (
	"fmt"
)

func Sortdata(data1 []int, data2 []int) {
	for i := 0; i < len(data2); i++ {
		data1 = append(data1, data2[i])
		for j := len(data1) - 1; j > 0; j-- {
			if data1[j] < data1[j-1] {
				data1[j], data1[j-1] = data1[j-1], data1[j]
			}
		}
	}
	fmt.Println(data1)
}

// 双指针从后往前
func merge(data1 []int, m int, data2 []int, n int) []int {

	i := m - 1
	j := n - 1
	for i := 0; i < len(data2); i++ {
		data1 = append(data1, 0)
	}
	p := len(data1) - 1
	for i >= 0 && j >= 0 {
		if data1[i] >= data2[j] {
			data1[p] = data1[i]
			i--
			p--
		} else if data1[i] <= data2[j] {
			data1[p] = data2[j]
			p--
			j--
		}
	}
	return data1
}
func main() {
	data1 := []int{1, 3, 9, 10}
	data2 := []int{2, 4, 11, 12}
	merge(data1, 4, data2, 4)
}
