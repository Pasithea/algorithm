"""
    Selection sort.
"""


class SelectionSort:
    def sort(self, l: list) -> list:
        length = len(l)
        if length <= 1: return l
        for i in range(length):
            min_val = l[i]
            min_ix = i
            for j in range(i + 1, length):
                if l[j] < min_val:
                    min_val = l[j]
                    min_ix = j
            l[i], l[min_ix] = l[min_ix], l[i]
        return l


if __name__ == '__main__':
    sel = SelectionSort()
    l = [3, 9, 6, 4, 7, 2]
    print(sel.sort(l))
