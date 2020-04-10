// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/


// Definition for a Node.
type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

// time O(N) space O(N)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	arr := make([]*Node, 0, 4096)
	arr = append(arr, root)
	for len(arr) > 0 {
		size := len(arr)
		for i := 0; i < size; i++ {
			x := arr[0]
			arr = arr[1:]
			if i < size - 1 {
				x.Next = arr[0]
			}
			if x.Left != nil {
				arr = append(arr, x.Left)
			}
			if x.Right != nil {
				arr = append(arr, x.Right)
			}
		}
	}
	return root
}

// // time O(N) space O(1)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftmost = leftmost.Left
	}
	return root
}
