package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) (t *TreeNode) {
	if len(nums) == 0 {
		return
	}
	// 找到左右子树并定位根节点
	mid := len(nums) / 2
	root := nums[mid]
	t = &TreeNode{root, nil, nil}
	if len(nums[:mid]) > 0 {
		t.Left = sortedArrayToBST(nums[:mid])
	}
	if len(nums[mid+1:]) > 0 {
		t.Right = sortedArrayToBST(nums[mid+1:])
	}
	return t
}

// 遍历二叉树
func read(root *TreeNode) {
	if root == nil {
		return
	}
	read(root.Left)
	fmt.Println(root.Val)
	read(root.Right)
}
func main() {
	data := []int{5, 6, 7, 8, 10, 11, 13, 14}
	root := sortedArrayToBST(data)
	read(root)
}
