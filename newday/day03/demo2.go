package main

import (
	"fmt"
)

// 插入排序
func InsertSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

func main() {
	data1 := []int{0, 1, 2, 0, 9, 1, 2, 3}
	data2 := []int{}
	data3 := []int{1}
	data4 := []int{9, 0}
	fmt.Println(InsertSort(data1))
	fmt.Println(InsertSort(data2))
	fmt.Println(InsertSort(data3))
	fmt.Println(InsertSort(data4))
}
