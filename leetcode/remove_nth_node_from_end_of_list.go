// https://leetcode.com/problems/remove-nth-node-from-end-of-list/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	if p2 == nil {
		return head.Next
	}
	p2 = p2.Next
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p1.Next = p1.Next.Next
	return head
}
