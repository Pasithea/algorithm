"""
    Use linked list to check a string is palindrome or not.
"""


class Node:
    def __init__(self, value=None, _next=None):
        self.value = value
        self.next = _next


class Solution:
    def __init__(self, head: Node):
         self.head = head

    def __call__(self) -> bool:
        if self.head is None or self.head.next is None:
            return True
        prev = None
        slow = self.head
        fast = self.head
        while fast is not None and fast.next is not None:
            fast = fast.next.next
            next_ = slow.next
            slow.next = prev
            prev = slow
            slow = next_
        if fast is not None:
            slow = slow.next
        while slow is not None:
            if slow.value != prev.value:
                return False
            slow = slow.next
            prev = prev.next
        return True


def create_linked_list(s: str) -> Node:
    head = Node(s[0])
    pre = head
    for v in s[1:]:
        n = Node(v)
        pre.next = n
        pre = n
    return head


if __name__ == '__main__':
    assert Solution(create_linked_list('sator arepo tenet opera rotas'))() is True
