"""
    Quick sort.
"""


class QuickSort:
    def sort(self, l: list):
        self.qsort(l, 0, len(l)-1)

    def qsort(self, l: list, low: int, high: int):
        if low >= high: return
        m = self.partition(l, low, high)
        self.qsort(l, low, m-1)
        self.qsort(l, m+1, high)

    def partition(self, l: list, low: int, high: int):
        pivot, i = l[high], low
        for j in range(low, high):
            if l[j] < pivot:
                l[i], l[j] = l[j], l[i]
                i += 1
        l[i], l[high] = l[high], l[i]
        return i


if __name__ == '__main__':
    qs = QuickSort()
    ll = [4, 9, 1, 7, 3, 2, 8, 0, 6, 5]
    qs.sort(ll)
    print(ll)
