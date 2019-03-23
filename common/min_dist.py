"""
    Min distance.
"""


MAP = [[1, 3, 5, 9],
       [2, 1, 3, 4],
       [5, 2, 6, 7],
       [6, 8, 4, 3]]


class MinDist:
    def __init__(self):
        self.n = len(MAP)
        self.min_dist = float('inf')

    def back_tracing(self, i: int, j: int, dist: int) -> None:
        if i == (self.n - 1) and j == (self.n - 1):
            fin_dist = dist + MAP[self.n-1][self.n-1]
            if fin_dist < self.min_dist:
                self.min_dist = fin_dist
                return
        if i < (self.n - 1):
            self.back_tracing(i+1, j, dist+MAP[i][j])
        if j < (self.n - 1):
            self.back_tracing(i, j+1, dist+MAP[i][j])

    def dp1(self) -> int:
        states = [[0 for _ in range(self.n)] for _ in range(self.n)]
        sum_ = 0
        for i in range(self.n):
            sum_ += MAP[i][0]
            states[i][0] = sum_
        sum_ = 0
        for j in range(self.n):
            sum_ += MAP[0][j]
            states[0][j] = sum_
        for i in range(1, self.n):
            for j in range(1, self.n):
                states[i][j] = min(states[i-1][j], states[i][j-1]) + MAP[i][j]
        return states[self.n-1][self.n-1]

    def dp2(self) -> int:
        mem = [[0 for _ in range(self.n)] for _ in range(self.n)]
        def min_dist(i: int, j: int):
            if i == 0 and j == 0:
                return MAP[0][0]
            if mem[i][j] > 0:
                return mem[i][j]
            min_left = float('inf')
            if j >= 1:
                min_left = min_dist(i, j-1)
            min_up = float('inf')
            if i >= 1:
                min_up = min_dist(i-1, j)
            cur_min = min(min_left, min_up) + MAP[i][j]
            mem[i][j] = cur_min
            return cur_min
        return min_dist(self.n-1, self.n-1)


if __name__ == '__main__':
    min_dist = MinDist()
    min_dist.back_tracing(0, 0, 0)
    print(min_dist.min_dist)
    print(min_dist.dp1())
    print(min_dist.dp2())
