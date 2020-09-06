package main

import (
	"fmt"
)

var preOrderTraversal, inOrderTraversal []int

// BinaryTreeNode .
type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func main() {
	preOrderTraversal = append(preOrderTraversal, 1, 2, 4, 7, 3, 5, 6, 8)
	inOrderTraversal = append(inOrderTraversal, 4, 7, 2, 1, 5, 3, 8, 6)
	length := len(preOrderTraversal)
	root := rebuildTree(preOrderTraversal, inOrderTraversal, length)
	printPreOrder(root)
}

func rebuildTree(pt, it []int, length int) *BinaryTreeNode {
	rv := pt[0]
	root := &BinaryTreeNode{rv, nil, nil}
	var i, v int
	for i, v = range it {
		if v == rv {
			break
		}
	}
	if i < length {
		leftit := it[:i]
		rightit := it[i+1:]
		leftpt := pt[1 : i+1]
		rightpt := pt[i+1:]

		if len(leftit) > 0 && len(leftpt) > 0 {
			root.left = rebuildTree(leftpt, leftit, len(leftit))
		}
		if len(rightit) > 0 && len(rightpt) > 0 {
			root.right = rebuildTree(rightpt, rightit, len(rightit))
		}
	}
	return root
}

func printPreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Println(root.value)
	}
	if root.left != nil {
		printPreOrder(root.left)
	}
	if root.right != nil {
		printPreOrder(root.right)
	}
}
