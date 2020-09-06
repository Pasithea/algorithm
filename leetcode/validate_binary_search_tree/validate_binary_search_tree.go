// https://leetcode.com/problems/validate-binary-search-tree/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	maxInt := 1<<63 - 1
	minInt := -1 << 63
	return isValid(root, maxInt, minInt)
}

func isValid(root *TreeNode, maxInt, minInt int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxInt || root.Val <= minInt {
		return false
	}
	return isValid(root.Left, root.Val, minInt) && isValid(root.Right, maxInt, root.Val)
}
