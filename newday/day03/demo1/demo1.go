package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumy := new(ListNode)
	dumy.Next = head
	first := dumy
	second := dumy
	for i := 1; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dumy.Next
}
func main() {

}
