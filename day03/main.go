/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-24 17:08:53
 * @LastEditTime: 2019-08-24 17:26:46
 * @LastEditors: Please set LastEditors
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	x, y := 0, 0
	p := l1
	q := l2
	curr := dummyHead
	carry := 0
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}

	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummyHead.Next
}
func main() {

}
