"""
    Counting inversions using merge sort.
"""

from typing import List


class InversionCount:
    def count(self, l: List[int]) -> int:
        n = len(l)
        self._num = 0
        self.merge_sort_count(l, 0, n-1)
        return self._num

    def merge_sort_count(self, l: List[int], low: int, high: int) -> None:
        if low >= high: return
        mid = (low + high) // 2
        self.merge_sort_count(l, low, mid)
        self.merge_sort_count(l, mid+1, high)
        self.merge(l, low, mid, high)

    def merge(self, l:List[int], low: int, mid: int, high: int) -> None:
        i, j = low, mid+1
        tmp = list()
        while i <= mid and j <= high:
            if l[i] <= l[j]:
                tmp.append(l[i])
                i += 1
            else:
                self._num += (mid - i + 1) # greater than l[j] element count between i and mid.
                tmp.append(l[j])
                j += 1
        start = i if i <= mid else j
        end = mid if i <= mid else high
        tmp.extend(l[start: end+1])
        l[low: high+1] = tmp


if __name__ == '__main__':
    ic = InversionCount()
    l = [4, 9, 1, 7, 3, 2, 8, 6, 5]
    n = ic.count(l)
    print(n)
