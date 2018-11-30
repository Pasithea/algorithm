"""
    Insertion sort.
"""


class InsertionSort:
    def sort(self, l: list) -> list:
        length = len(l)
        if length <= 1: return l
        for i in range(1, length):
            val = l[i]
            j = i - 1
            while j >= 0:
                if l[j] > val:
                    l[j+1] = l[j]
                else:
                    break
                j -= 1
            l[j+1] = val
        return l


if __name__ == '__main__':
    ins = InsertionSort()
    l = [4, 7, 2, 1, 3, 9]
    # \\ also you use array instead list
    # from array import array
    # l = array('b', (4, 7, 2, 1, 3, 9))
    print(ins.sort(l))
