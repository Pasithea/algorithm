"""
    Implement LRU with doubly linked list.
"""

from typing import Optional


class Node:
    def __init__(self, value=None, _next=None, _prev=None):
        self.value = value
        self.next = _next
        self.prev = _prev

    def __repr__(self):
        prev_ = self.prev.value if self.prev else None
        next_ = self.next.value if self.next else None
        return '<Node prev: %s, self: %s, next: %s>' % (prev_, self.value, next_)


class LRUCache:
    def __init__(self, max_size=5):
        self._max_size = max_size
        self._head = None
        self._size = 0

    def find(self, val: Optional) -> Node:
        if self._head is None:
            self._head = Node(val)
            self._size += 1
            return self._head
        cur = self._head
        end = None
        while cur is not None:
            nxt, prv = cur.next, cur.prev
            if cur.value == val:
                if cur is not self._head:
                    prv.next = nxt
                    if nxt is not None:
                        nxt.prev = prv
                    cur.next = self._head
                    cur.prev = None
                    self._head.prev = cur
                    self._head = cur
                return self._head
            end = cur
            cur = nxt
        if self._size < self._max_size:
            self._size += 1
        else:
            prv = end.prev
            prv.next = None
            end.prev = None
            end.next = None
        n = Node(val)
        self._head.prev = n
        n.next = self._head
        self._head = n
        return self._head

    def clear(self):
        self._head = None
        self._size = 0

    def __getitem__(self, item: int) -> Node:
        cur = self._head
        for _ in range(item):
            try:
                cur = cur.next
            except AttributeError:
                raise IndexError('linked list out of range')
        if cur is None:
            raise IndexError('linked list out of range')
        return cur

    def __repr__(self) -> str:
        stl = []
        cur = self._head
        for _ in range(self._size):
            if cur is None:
                continue
            stl.append(str(cur.value))
            cur = cur.next
        return ' -> '.join(stl) if stl else 'empty'


def main(size: int):
    cache = LRUCache(size)
    import random
    import string
    for _ in range(size * 10):
        cache.find(random.choice(string.ascii_uppercase))
    return cache


if __name__ == '__main__':
    print(main(10))
