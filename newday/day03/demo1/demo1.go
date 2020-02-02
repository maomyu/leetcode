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

// 删除倒数第K个结点
func removeNthFromEnd(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	first := head
	second := head
	for i := 0; i < k+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
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
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	next := createLink(data, 0)
	head := new(ListNode)
	head.Next = next
	readLink(head.Next)
	removeNthFromEnd(head, 0)
	readLink(head.Next)
}
