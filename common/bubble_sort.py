"""
    Bubble Sort.
"""


class BubbleSort:
    def sort(self, l: list) -> list:
        length = len(l)
        if length <= 1: return l
        for i in range(length):
            flag = False
            for j in range(length-1):
                if l[j] > l[j+1]:
                    l[j], l[j+1] = l[j+1], l[j]
                    flag = True
            if not flag:
                break


if __name__ == '__main__':
    bs = BubbleSort()
    l = [4, 1, 6, 7, 9, 3]
    bs.sort(l)
    print(l)
