"""
https://leetcode.com/problems/two-sum/
"""


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hashmap = dict()
        for index, value in enumerate(nums):
            hashmap[value] = index
        for item in hashmap:
            other = target - item
            if other in hashmap and other != item:
                return [hashmap[item], hashmap[other]]
            if other == item:
                other_index = nums.index(other)
                if other_index != hashmap[item]:
                    return [other_index, hashmap[item]]
