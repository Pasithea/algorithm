// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		rv    []int
		stack []*TreeNode
	)
	if root == nil {
		return rv
	}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rv = append(rv, cur.Val)
		cur = cur.Right
	}
	return rv
}
