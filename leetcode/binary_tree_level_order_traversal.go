// https://leetcode.com/problems/binary-tree-level-order-traversal/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	return appendLevel(root, arr, 0)
}

func appendLevel(node *TreeNode, arr [][]int, level int) [][]int {
	if node == nil {
		return arr
	}
	if len(arr) < level + 1 {
		arr = append(arr, []int{})
	}
	arr[level] = append(arr[level], node.Val)
	arr = appendLevel(node.Left, arr, level+1)
	arr = appendLevel(node.Right, arr, level+1)
	return arr
}
