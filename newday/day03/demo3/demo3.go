package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个链表
func creatLink(data []int,index int)(head *ListNode){
	if index >=len(data){
		return head
	}
	head = &ListNode{data[index],nil}
	head.Next = creatLink(data,index+1)
	return head
}

// 判断是否是回文链表,方法一，复制数组
func isPalindrome(head *ListNode) bool {
	p :=head
	data :=[]int{}
	for p!=nil{
		data = append(data,p.Val)
		p = p.Next
	}
	i,j :=0,len(data)-1
	for i<=j{
		if data[i] != data[j]{
			return false
		}
		i++
		j--
	}
	return true
}


func main() {
	data :=[]int{1,2,3,4,3,2,1}
	head :=creatLink(data,0)
	fmt.Println(isPalindrome(head))
}
