"""
    Shell sort.
"""


class ShellSort:
    def sort(self, l: list) -> list:
        length = len(l)
        if length <= 1: return l
        gap = length // 2
        while gap > 0:
            for i in range(gap, length):
                tmp = l[i]
                j = i
                while j >= gap and l[j - gap] > tmp:
                    l[j] = l[j - gap]
                    j -= gap
                l[j] = tmp
            gap = gap // 2
        return l


if __name__ == '__main__':
    she = ShellSort()
    l = [4, 9, 1, 0, 3, 8, 2, 6, 5, 7]
    print(she.sort(l))
