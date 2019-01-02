"""
    Counting sort.
"""


class CountingSort:
    def sort(self, l: list):
        length = len(l)
        if length <= 1: return
        counts = [0] * (max(l) + 1)
        for i in l:
            counts[i] += 1
        for j in range(1, length):
            counts[j] = counts[j-1] + counts[j]
        new_sort = [0] * length
        for i in l:
            index = counts[i] - 1
            new_sort[index] = i
            counts[i] -= 1
        l[:] = new_sort[:]


if __name__ == '__main__':
    cs = CountingSort()
    l = [4, 9, 1, 7, 3, 2, 8, 0, 6, 5]
    cs.sort(l)
    print(l)
