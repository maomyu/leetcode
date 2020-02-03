package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个链表
func createLink(data []int, index int) (head *ListNode) {
	if len(data) <= index {
		return
	}
	head = &ListNode{data[index], nil}
	head.Next = createLink(data, index+1)
	return head
}

// 合并两个有序的链表
func MeargeList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MeargeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = MeargeList(l1, l2.Next)
		return l2
	}
}

// 和并两个有序数组
func Merage(data1 []int, data2 []int) {
	i := len(data1) - 1
	for i := 0; i < len(data2); i++ {
		data1 = append(data1, data2[i])
	}
	j := len(data1) - 1
	for i <= j && i >= 0 && j >= 0 {
		if data1[i] <= data1[j] {
			j--
		} else if data1[i] > data1[j] {
			data1[i], data1[j] = data1[j], data1[i]
			i--
		}
	}
	fmt.Println(data1)
}

// 获取链表数据
func readLink(head *ListNode) {
	if head == nil {
		fmt.Println()
		return
	}
	fmt.Print(head.Val, " ")
	readLink(head.Next)
}
func main() {
	data1 := []int{3, 3, 5, 7}
	data2 := []int{1, 2, 9, 11}
	l1 := createLink(data1, 0)
	l2 := createLink(data2, 0)
	l := MeargeList(l1, l2)
	readLink(l)
}
