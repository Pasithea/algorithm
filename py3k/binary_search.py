"""
    Binary search.
"""

from typing import Union


class BinarySearch:
    def bsearch(self, sorted_list: list, target: int) -> Union[int, None]:
        low, high = 0, len(sorted_list) - 1
        while low <= high:
            mid = low + (high - low) // 2
            if sorted_list[mid] == target:
                return mid
            elif sorted_list[mid] < target:
                low = mid + 1
            else:
                high = mid - 1
        return None


if __name__ == '__main__':
    bs = BinarySearch()
    sorted_list = [1, 3, 5, 7, 9, 10, 20, 40, 60, 80, 100]
    for i in reversed(sorted_list):
        print(bs.bsearch(sorted_list, i))
