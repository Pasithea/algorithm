"""
    0-1 Bag.
"""


from typing import List


class Bag01:
    """
    back track.
    """

    def __init__(self, items: List[int], w: int):
        self.max_w = float('-inf')
        self.items = items
        self.n = len(items)
        self.w = w

    def cal(self, i: int, cw: int) -> None:
        # cw: current weight
        if cw == self.w or i == self.n:
            if cw > self.max_w:
                self.max_w = cw
            return
        self.cal(i + 1, cw)
        if cw + self.items[i] <= self.w:
            self.cal(i + 1, cw + self.items[i])


class Bag01knapsack:
    """
    DP.
    """

    def __init__(self, items: List[int], w: int):
        self.items = items
        self.n = len(items)
        self.w = w
        self.states = [[False for _ in range(w+1)] for _ in range(self.n)]

    def cal(self) -> int:
        self.states[0][0] = True
        self.states[0][self.items[0]] = True
        for i in range(1, self.n):
            for j in range(self.w + 1):
                if self.states[i-1][j]:
                    self.states[i][j] = True
            j = 0
            while j <= (self.w - self.items[i]):
                if self.states[i-1][j]:
                    self.states[i][j+self.items[i]] = True
                j += 1
        for x in range(self.w, -1, -1):
            if self.states[self.n-1][x]:
                return x
        return 0


class Bag01DP:
    """
    DP with item value.
    """

    def __init__(self, weight: List[int], value: List[int], w: int):
        self.weight = weight
        self.value = value
        self.n = len(self.weight)
        self.w = w
        self.states = [[-1 for _ in range(w+1)] for _ in range(self.n)]

    def cal(self) -> int:
        self.states[0][0] = 0
        self.states[0][self.weight[0]] = self.value[0]
        for i in range(1, self.n):
            for j in range(self.w + 1):
                if self.states[i-1][j] != -1:
                    self.states[i][j] = self.states[i-1][j]
            j = 0
            while j <= (self.w - self.weight[i]):
                if self.states[i-1][j] >= 0:
                    v = self.states[i-1][j] + self.value[i]
                    if v > self.states[i][j+self.weight[i]]:
                        self.states[i][j+self.weight[i]] = v
                j += 1

        max_v = -1
        for x in range(self.w, -1, -1):
            if self.states[self.n-1][x] > max_v:
                max_v = self.states[self.n-1][x]
        return max_v


if __name__ == '__main__':
    items = [3, 5, 10, 13, 20]
    w = 42
    bag = Bag01(items, w)
    bag.cal(0, 0)
    print('max weight: %d' % bag.max_w)
    bag_knapsack = Bag01knapsack(items, w)
    print('max weight knapsack:', bag_knapsack.cal())
    values = [8, 4, 3, 7, 10]
    bagdp = Bag01DP(items, values, w)
    print('max value:', bagdp.cal())
    print(bagdp.states)
