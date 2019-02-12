"""
https://leetcode.com/problems/reverse-integer/
"""


class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        n = 1 if x >= 0 else -1
        s = str(abs(x))
        result = int(''.join(reversed(s))) * n
        # or result = int(s[::-1]) * n
        # or result = int(''.join(s[x] for x in range(len(s)-1, -1, -1))) * n
        if result < (-2 ** 31) or result > (2 ** 31):
            return 0
        return result
