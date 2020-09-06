// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}

}

// if not bst, the following is the common way that use heap.
//func kthSmallest(root *TreeNode, k int) int {
//	heap := make([]int, 0)
//	arr := make([]*TreeNode, 0)
//	arr = append(arr, root)
//	for len(arr) != 0 {
//		size := len(arr)
//		for i := 0; i < size; i++ {
//			node := arr[0]
//			arr = arr[1:]
//			if node.Left != nil {
//				arr = append(arr, node.Left)
//			}
//			if node.Right != nil {
//				arr = append(arr, node.Right)
//			}
//			if len(heap) < k {
//				heap = append(heap, node.Val)
//				up(heap, len(heap) - 1)
//			} else {
//				if node.Val < heap[0] {
//					heap[0] = node.Val
//					down(heap, 0, len(heap) - 1)
//				}
//			}
//		}
//	}
//	return heap[0]
//}

func up(heap []int, n int) {
	for n > 0 {
		p := (n - 1) / 2
		if p == n {
			return
		}
		if heap[n] > heap[p] {
			heap[p], heap[n] = heap[n], heap[p]
		} else {
			return
		}
		n = p
	}
}

func down(heap []int, n, k int) {
	for n <= k {
		c := n
		c1 := n*2 + 1
		c2 := c1 + 1
		if c1 <= k && heap[c1] > heap[c] {
			c = c1
		}
		if c2 <= k && heap[c2] > heap[c] {
			c = c2
		}
		if c == n {
			return
		}
		heap[n], heap[c] = heap[c], heap[n]
		n = c
	}
}
