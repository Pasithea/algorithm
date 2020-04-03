// https://leetcode.com/problems/add-two-numbers/


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		cur  *ListNode
		head *ListNode
	)
	carry := 0
	head = new(ListNode)
	cur = head
	for l1 != nil || l2 != nil || carry != 0 {
		var (
			t1 int
			t2 int
		)
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}
		if cur.Next == nil {
			cur.Next = new(ListNode)
		}
		cur.Next.Val = t1 + t2 + carry
		carry = cur.Next.Val / 10
		cur.Next.Val %= 10
		cur = cur.Next
	}
	return head.Next
}
