"""
    Max-heap.
    Python standard library heapq is a min-heap implementation.
"""

from typing import Union


class Heap:
    def __init__(self, capacity: int=10):
        self._heap = []
        self._capacity = capacity
        self.count = 0

    def push(self, value: int) -> None:
        if self.count >= self._capacity:
            return
        self._heap.append(value)
        self.count += 1
        pos = len(self._heap) - 1
        while (pos - 1) >> 1 >= 0 and self._heap[pos] > self._heap[(pos - 1) >> 1]:
            self._heap[pos], self._heap[(pos - 1) >> 1] = self._heap[(pos - 1) >> 1], self._heap[pos]
            pos = (pos - 1) >> 1

    def poptop(self) -> Union[None, int]:
        if len(self._heap) == 0:
            return
        last = self._heap.pop()
        self.count -= 1
        if len(self._heap) == 0:
            v = last
        else:
            v = self._heap[0]
            self._heap[0] = last
        pos = 0
        end_pos = len(self._heap) - 1
        while True:
            max_pos = pos
            if (pos << 1) + 1 <= end_pos and self._heap[pos] < self._heap[(pos << 1) + 1]:
                max_pos = (pos << 1) + 1
            if (pos << 1) + 2 <= end_pos and self._heap[max_pos] < self._heap[(pos << 1) + 2]:
                max_pos = (pos << 1) + 2
            if max_pos == pos:
                break
            self._heap[pos], self._heap[max_pos] = self._heap[max_pos], self._heap[pos]
            pos = max_pos
        return v

    def __repr__(self):
        return '<Heap> %s' % (str(self._heap))


if __name__ == '__main__':
    h = Heap()
    import random
    l = list(range(10))
    random.shuffle(l)
    print(l)
    for i in l:
        h.push(i)
    print(h._heap)
    for _ in range(h.count):
        print(h.poptop())
