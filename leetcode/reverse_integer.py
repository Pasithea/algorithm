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
        x *= n
        rv = 0
        while x != 0:
            rv = rv * 10 + x % 10
            x //= 10
        rv *= n
        if rv < (-2 ** 31) or rv > (2 ** 31):
            return 0
        return rv
