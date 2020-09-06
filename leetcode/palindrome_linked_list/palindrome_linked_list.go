// https://leetcode.com/problems/palindrome-linked-list/
package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// reverse
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	if fast != nil {
		slow = slow.Next
	}
	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}
