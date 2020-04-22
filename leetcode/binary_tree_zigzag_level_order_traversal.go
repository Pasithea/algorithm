// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/


// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var rs [][]int
	rs = sub(root, 0, rs)
	for i := range rs {
		if i % 2 == 1 {
			reverse(rs[i])
		}
	}
	return rs
}

func sub(node *TreeNode, level int, rs [][]int) [][]int {
	if node == nil {
		return rs
	}
	if len(rs) <= level {
		rs = append(rs, []int{})
	}
	rs[level] = append(rs[level], node.Val)
	rs = sub(node.Left, level+1, rs)
	rs = sub(node.Right, level+1, rs)
	return rs
}

func reverse(s []int) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
