"""
https://leetcode.com/problems/first-bad-version/
"""


# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        b = g = 0
        start, end = 1, n
        while (b-g) != 1:
            m = (start + end) // 2
            if isBadVersion(m):
                b = m
                end = m
            else:
                g = m
                start = m + 1
        return b
