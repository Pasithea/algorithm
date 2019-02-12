"""
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
"""


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        hashmap = dict()
        if head is None:
            return
        hashmap[head.val] = 1
        p = head
        cur = p.next
        while cur is not None:
            if cur.val in hashmap:
                cur = cur.next
                p.next = cur
            else:
                hashmap[cur.val] = 1
                p = p.next
                cur = cur.next
        return head
