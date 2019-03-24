"""
    N queens puzzle.
"""


class NQueens:
    def __init__(self, size=8):
        self._size = size
        self._solotions = 0
        self._result = [-1] * self._size
        self.calculate(0)

    def calculate(self, row: int):
        if row == self._size:
            self._solotions += 1
            print('No.', self._solotions)
            self._print_queens()
            print(self._result)
            print('\n')
            return
        for column in range(self._size):
            if self._check_place(row, column):
                self._result[row] = column
                self.calculate(row+1)

    def _check_place(self, row: int, column: int):
        leftup = column - 1
        rightup = column + 1
        i = row - 1
        while i >= 0:
            if self._result[i] == column:
                return False
            if leftup >= 0:
                if self._result[i] == leftup:
                    return False
            if rightup < self._size:
                if self._result[i] == rightup:
                    return False
            leftup -= 1
            rightup += 1
            i -= 1
        return True

    def _print_queens(self):
        for row in range(self._size):
            line = ''
            for column in range(self._size):
                if self._result[row] == column:
                    line += 'Q '
                else:
                    line += '* '
            print(line)


if __name__ == '__main__':
    NQueens(size=8)
