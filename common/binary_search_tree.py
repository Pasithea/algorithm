"""
    Binary search tree.
"""

from typing import Union


class Node:
    def __init__(self, value: int):
        self.value = value
        self.left = None
        self.right = None


class BinarySearchTree:
    def __init__(self):
        self._root = None

    def find(self, value: int) -> Union[Node, None]:
        p = self._root
        while p is not None:
            if value < p.value:
                p = p.left
            elif value > p.value:
                p = p.right
            else:
                return p
        return None

    def insert(self, value: int) -> None:
        if self._root is None:
            self._root = Node(value)
            return
        p = self._root
        while p is not None:
            if value > p.value:
                if p.right is None:
                    p.right = Node(value)
                    return
                p = p.right
            else:
                if p.left is None:
                    p.left = Node(value)
                    return
                p = p.left

    def delete(self, value: int) -> None:
        p = self._root
        pp = None
        while p is not None and p.value != value:
            pp = p
            if value > p.value:
                p = p.right
            else:
                p = p.left
        if p is None:
            return

        if p.left is not None and p.right is not None:
            min_p = p.right
            min_pp = p
            while p.left is not None:
                min_p = p.left
                min_pp = p
            p.value = min_p.value
            p = min_p
            pp = min_pp

        if p.left is not None:
            child = p.left
        elif p.right is not None:
            child = p.right
        else:
            child = None

        if pp is None:
            self._root = child
        elif pp.left == p:
            pp.left = child
        else:
            pp.right = child
