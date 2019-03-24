"""
    Merge sort.
"""


class MergeSort:
    def sort(self, l: list):
        low = 0
        high = len(l) - 1
        self.sort_c(l, low, high)

    def sort_c(self, l: list, low: int, high: int):
        if low >= high: return
        mid = (low + high) // 2
        self.sort_c(l, low, mid)
        self.sort_c(l, mid+1, high)
        self.merge(l, low, mid, high)

    def merge(self, l: list, low: int, mid: int, high: int):
        i, j = low, mid+1
        tmp = list()
        while i <= mid and j <= high:
            if l[i] <= l[j]:
                tmp.append(l[i])
                i += 1
            else:
                tmp.append(l[j])
                j += 1
        start = i if i <= mid else j
        end = mid if i <= mid else high
        tmp.extend(l[start: end+1])
        l[low:high+1] = tmp


if __name__ == '__main__':
    mer = MergeSort()
    l = [4, 9, 1, 7, 3, 2, 8, 0, 6, 5]
    mer.sort(l)
    print(l)
