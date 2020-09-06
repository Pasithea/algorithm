// https://leetcode.com/problems/binary-tree-preorder-traversal/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var (
		rv    []int
		stack []*TreeNode
	)
	if root == nil {
		return rv
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root != nil {
			rv = append(rv, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return rv
}
