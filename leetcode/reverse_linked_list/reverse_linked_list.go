// https://leetcode.com/problems/reverse-linked-list/
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	p1.Next = nil
	for p2 != nil {
		next := p2.Next
		p2.Next = p1
		p1 = p2
		if next == nil {
			break
		}
		p2 = next
	}
	return p2
}
