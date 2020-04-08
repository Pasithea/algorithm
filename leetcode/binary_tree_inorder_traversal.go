// https://leetcode.com/problems/binary-tree-inorder-traversal/


// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		rs []int
	)
	rs = sub(root, rs)
	return rs
}

func sub (node *TreeNode, rs []int) []int {
	if node == nil {
		return rs
	}
	rs = sub(node.Left, rs)
	rs = append(rs, node.Val)
	rs = sub(node.Right, rs)
	return rs
}
