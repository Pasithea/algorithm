// https://leetcode.com/problems/intersection-of-two-linked-lists/
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		a *ListNode = headA
		b *ListNode = headB
		s *ListNode
		x *ListNode
	)
	for a != nil && b != nil {
		a = a.Next
		b = b.Next
	}
	if a == nil {
		s = headB
		x = headA
	} else {
		s = headA
		x = headB
	}
	for a != nil {
		a = a.Next
		s = s.Next
	}
	for b != nil {
		b = b.Next
		s = s.Next
	}
	for x != nil && s != nil {
		if x == s {
			return s
		}
		x = x.Next
		s = s.Next
	}
	return nil
}
