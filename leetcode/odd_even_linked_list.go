// https://leetcode.com/problems/odd-even-linked-list/


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := even
	for {
		if even != nil && even.Next != nil {
			odd.Next = even.Next
			odd = even.Next
			even.Next = odd.Next
			even = even.Next
		} else {
			break
		}
	}
	odd.Next = evenHead
	return head
}
