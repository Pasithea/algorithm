// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/


// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var (
		i int
	)
	for i = 0; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			break
		}
	}
	if i < len(inorder) {
		preorderLeft := preorder[1:i+1]
		preorderRight := preorder[i+1:]
		inorderLeft := inorder[:i]
		inorderRight := inorder[i+1:]
		root.Left = buildTree(preorderLeft, inorderLeft)
		root.Right = buildTree(preorderRight, inorderRight)
	}
	return root
}
