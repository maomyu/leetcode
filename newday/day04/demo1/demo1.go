package main

// 二叉树的最大的深度
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//

func maxDepth(root *TreeNode) (h int) {
	if root == nil {
		return
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left
	}
	return right
}
func main() {

}
