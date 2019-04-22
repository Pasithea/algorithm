// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildTree(nums, start, mid-1)
	node.Right = buildTree(nums, mid+1, end)
	return node
}
