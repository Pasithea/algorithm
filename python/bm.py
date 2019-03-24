"""
    Boyer–Moore string-search algorithm.
"""

from typing import List, Tuple


_SIZE = 256


class BM:
    def _generate_bc(self, pattern: str) -> List[int]:
        # bad character.
        bc = [-1 for _ in range(_SIZE)]
        for i, s in enumerate(pattern):
            bc[ord(s)] = i
        return bc

    def _generate_gs(self, pattern: str) -> Tuple[List[bool], List[int]]:
        # good suffix.
        m = len(pattern)
        suffix = [-1 for _ in range(m)]
        prefix = [False for _ in range(m)]
        for i in range(m - 1):
            j = i
            k = 0
            while j >= 0 and pattern[j] == pattern[m-1-k]:
                j -= 1
                k += 1
                suffix[k] = j + 1
            if j == -1:
                prefix[k] = True
        return prefix, suffix

    def _move_by_gs(self, bc_index: int, m: int, suffix: List[int], prefix: List[bool]):
        k = m - 1 - bc_index
        if suffix[k] != -1:
            return bc_index - suffix[k] + 1
        for r in range(bc_index+2, m):
            if prefix[m-r] is True:
                return r
        return m

    def bm(self, main: str, pattern: str) -> int:
        n = len(main)
        m = len(pattern)
        if n < m:
            return -1
        bc = self._generate_bc(pattern)
        prefix, suffix = self._generate_gs(pattern)
        i = 0
        while i <= (n - m):
            j = m - 1
            while j >= 0:
                if main[i+j] != pattern[j]:
                    break
                j -= 1
            if j < 0:
                return i
            x = j - bc[ord(main[i+j])]
            y = 0
            if j < (m - 1):
                y = self._move_by_gs(j, m, suffix, prefix)
            i = i + max(x, y)
        return -1


if __name__ == '__main__':
    b = BM()
    x = 'playplaystationxboxxbox'
    print(b.bm(x, '42'), b.bm(x, 'on'), b.bm(x, 'box'), b.bm(x, 'xx'))
